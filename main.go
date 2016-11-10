package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

type DBConfig struct {
	Database string
	User     string
	Password string
	Host     string
	Port     int
}

var queries []string

func main() {
	dbConfig := readCommandLineFlags()
	// connect to database
	db := connectDatabase(dbConfig)
	defer db.Close()
	// read yaml file
	filenames := flag.Args()
	if len(filenames) == 0 {
		panic("Please provide a yaml file to load")
	}
	filename := filenames[0]
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var data interface{}

	err = yaml.Unmarshal([]byte(input), &data)
	if err != nil {
		log.Fatal(err)
	}

	switch items := data.(type) {
	case map[interface{}]interface{}:
		for _, item := range items {
			createSQLQuery(item.(map[interface{}]interface{}))
		}
	case []interface{}:
		for _, item := range items {
			// fmt.Printf("%+v\n", item.(map[interface{}]interface{}))
			createSQLQuery(item.(map[interface{}]interface{}))
		}
	default:
		panic("cannot parse file")
	}
	commitQueries(db)
}

func readCommandLineFlags() *DBConfig {
	var database string
	var user string
	var password string
	var host string
	var port int
	flag.StringVar(&database, "database", "postgres", "database name")
	flag.StringVar(&user, "user", "postgres", "database user name")
	flag.StringVar(&password, "password", "", "databases password")
	flag.StringVar(&host, "host", "localhost", "database host")
	flag.IntVar(&port, "port", 5432, "database port")
	flag.Parse()

	return &DBConfig{
		Database: database,
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
	}
}

func connectDatabase(dbConfig *DBConfig) *sql.DB {
	params := []string{}
	params = append(params, fmt.Sprintf("dbname=%s", dbConfig.Database))
	params = append(params, fmt.Sprintf("user=%s", dbConfig.User))
	if dbConfig.Password != "" {
		params = append(params, fmt.Sprintf("password=%s", dbConfig.Password))
	}
	params = append(params, fmt.Sprintf("host=%s", dbConfig.Host))
	params = append(params, fmt.Sprintf("port=%d", dbConfig.Port))
	params = append(params, fmt.Sprintf("sslmode=disable"))
	connString := strings.Join(params, " ")
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Print(err)
		log.Fatal("Error: The data source arguments are not valid")
	}
	err = db.Ping()
	if err != nil {
		log.Print(err)
		log.Fatal("Error: Couldn't stablish connection with the database")
	}
	return db
}

func createSQLQuery(record map[interface{}]interface{}) error {
	var columns []string
	var values []string
	tableName := ""
	for col, val := range record {
		if col == "table_name" {
			tableName = val.(string)
			continue
		}
		columns = append(columns, col.(string))
		switch val.(type) {
		case string:
			v := fmt.Sprintf("'%s'", val.(string))
			values = append(values, v)
		case int:
			v := fmt.Sprintf("%d", val.(int))
			values = append(values, v)
		case time.Time:
			v := fmt.Sprintf("to_date(:%d, 'HH24:MI:SS')", val.(time.Time))
			values = append(values, v)
		default:
			values = append(values, val.(string))
		}
	}
	q := fmt.Sprintf(
		"INSERT INTO %s(%s) values(%s)",
		tableName,
		strings.Join(columns, ", "),
		strings.Join(values, ", "),
	)

	queries = append(queries, q)
	return nil
}

func commitQueries(db *sql.DB) {
	for _, q := range queries {
		fmt.Printf("%s\n", q)
		stmt, err := db.Prepare(q)
		if err != nil {
			log.Fatal(err)
		}
		_, err = stmt.Exec()
		if err != nil {
			log.Fatal(err)
		}
	}
}

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	filename := os.Args[1]
	dbConfig := readCommandLineFlags()
	// connect to database
	db := connectDatabase(dbConfig)
	defer db.Close()
	// read yaml file
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
	return &DBConfig{
		Database: *flag.String("database", "postgres", "database name"),
		User:     *flag.String("user", "postgres", "database user name"),
		Password: *flag.String("password", "", "databases password"),
		Host:     *flag.String("host", "localhost", "database host"),
		Port:     *flag.Int("port", 5432, "database port"),
	}
}

func connectDatabase(dbConfig *DBConfig) *sql.DB {
	connString := fmt.Sprintf(
		"user=%s dbname=%s host=%s port=%d password=%s sslmode=disable",
		dbConfig.User,
		dbConfig.Database,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Password,
	)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Error: The data source arguments are not valid")
	}
	err = db.Ping()
	if err != nil {
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

package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/objectizer/gofixtures/database"
	"github.com/objectizer/gofixtures/utils"
)

var queries []string

func main() {
	cli := &utils.CLI{}
	cli.ReadCommandLineFlags()
	dbConf, err := cli.DatabaseConf()
	if err != nil {
		log.Fatal(err)
	}
	// connect to database
	db := database.ConnectDatabase(dbConf)
	defer db.Close()

	files := cli.FilesToParse()
	queries := []string{}
	for _, f := range files {
		data, err := utils.ParseYAML(f)
		if err != nil {
			panic(err)
		}
		switch items := data.(type) {
			case map[interface{}]interface{}:
				for _, item := range items {
					q, err := utils.BuildQuery(item.(map[interface{}]interface{}))
					if err != nil {
						panic(err)
					}
					queries = append(queries, q)
				}
			case []interface{}:
				for _, item := range items {
					q, err := utils.BuildQuery(item.(map[interface{}]interface{}))
					if err != nil {
						panic(err)
					}
					queries = append(queries, q)
				}
			default:
				panic("cannot parse file")
		}
		fmt.Printf("Loading %s...\n", f)
	}

	commitQueries(queries, db)
	fmt.Println("Finished Successfully...")
}

func commitQueries(queries []string, db *sql.DB) {
	for _, q := range queries {
		stmt, err := db.Prepare(q)
		if err != nil {
			log.Printf("%s\n", q)
			log.Fatal(err)
		}
		_, err = stmt.Exec()
		if err != nil {
			log.Printf("%s\n", q)
			log.Fatal(err)
		}
	}
}

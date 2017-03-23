package main

import (
	"database/sql"
	"fmt"

	"github.com/emostafa/gofixtures/database"
	"github.com/emostafa/gofixtures/utils"
)

var queries []string

func main() {
	cli := &utils.CLI{}
	cli.ReadCommandLineFlags()
	dbConf, err := cli.DatabaseConf()
	if err != nil {
		fmt.Println(err)
		return
	}
	// connect to database
	db, err := database.ConnectDatabase(dbConf)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	files, err := cli.FilesToParse()
	if err != nil {
		fmt.Println(err)
		return
	}
	queries := []string{}
	for _, f := range files {
		data, err := utils.ParseYAML(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch items := data.(type) {
		case map[interface{}]interface{}:
			for _, item := range items {
				q, err := utils.BuildQuery(item.(map[interface{}]interface{}))
				if err != nil {
					fmt.Println(err)
					return
				}
				queries = append(queries, q)
			}
		case []interface{}:
			for _, item := range items {
				q, err := utils.BuildQuery(item.(map[interface{}]interface{}))
				if err != nil {
					fmt.Println(err)
					return
				}
				queries = append(queries, q)
			}
		default:
			fmt.Println("cannot parse file")
			return
		}
		fmt.Printf("Loading %s...\n", f)
	}

	err = commitQueries(queries, db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Finished Successfully...")
}

func commitQueries(queries []string, db *sql.DB) error {
	for _, q := range queries {
		stmt, err := db.Prepare(q)
		if err != nil {
			return err
		}
		_, err = stmt.Exec()
		if err != nil {
			return err
		}
	}
	return nil
}

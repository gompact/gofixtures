package main

import (
	"fmt"
)

var queries []string

func main() {
	cli := &CLI{}
	cli.ReadCommandLineFlags()
	dbConf, err := cli.DatabaseConf()
	if err != nil {
		fmt.Println(err)
		return
	}

	// connect to database
	db, err := ConnectDatabase(dbConf)
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
		data, err := ParseYAML(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch items := data.(type) {
		case map[interface{}]interface{}:
			for _, item := range items {
				q, err := BuildQuery(item.(map[interface{}]interface{}))
				if err != nil {
					fmt.Println(err)
					return
				}
				queries = append(queries, q)
			}
		case []interface{}:
			for _, item := range items {
				q, err := BuildQuery(item.(map[interface{}]interface{}))
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

	err = CommitQueries(queries, db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Finished Successfully...")
}

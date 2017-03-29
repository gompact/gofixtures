package main

import (
	"database/sql"
	"flag"
	"fmt"
	gf "github.com/emostafa/gofixtures"
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
	db, err := gf.ConnectDatabase(dbConf)
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
		data, err := gf.ParseYAML(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		// switch items := data.(type) {
		// case map[interface{}]interface{}:
		// 	for _, item := range items {
		// 		q, err := utils.BuildQuery(item.(map[interface{}]interface{}))
		// 		if err != nil {
		// 			fmt.Println(err)
		// 			return
		// 		}
		// 		queries = append(queries, q)
		// 	}
		// case []interface{}:
		// 	for _, item := range items {
		// 		q, err := utils.BuildQuery(item.(map[interface{}]interface{}))
		// 		if err != nil {
		// 			fmt.Println(err)
		// 			return
		// 		}
		// 		queries = append(queries, q)
		// 	}
		// default:
		// 	fmt.Println("cannot parse file")
		// 	return
		// }
		fmt.Printf("Loading %s...\n", f)
		for _, item := range items {
			q, err := gf.BuildQuery(item.(map[interface{}]interface{}))
			if err != nil {
				fmt.Println(err)
				return
			}
			queries = append(queries, q)
		}
	}

	err = gf.CommitQueries(queries, db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Finished Successfully...")
}

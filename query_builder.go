package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

//BuildQuery buils a SQL query for passed record
func BuildQuery(record map[interface{}]interface{}) (string, error) {
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

	return q, nil
}

func CommitQueries(queries []string, db *sql.DB) error {
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

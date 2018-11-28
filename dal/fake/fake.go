package fake

import (
	"fmt"
	"log"
	"strings"

	"github.com/ishehata/gofixtures/v3/entity"
)

// New Creates
func New(config entity.DBConfig) *FakeDatastore {
	return &FakeDatastore{}
}

// FakeDatastore resembles a fake db which just prints SQL statements
type FakeDatastore struct {
}

// Connect connects to postgresql db based on parameters of DBConfig object
func (datastore *FakeDatastore) Connect() error {
	return nil
}

func (datastore *postgresDatastore) createTable(tableName string, columns []string) error {
	columnsDef := strings.Join(columns, " VARCHAR, ")
	columnsDef += " VARCHAR"
	q := fmt.Sprintf("CREATE TABLE IF NOT EXISTS public.%s (%s)", tableName, columnsDef)
	fmt.Println(q)
	return err
}

func keys(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func (datastore *postgresDatastore) Insert(fixture entity.Fixture) error {
	if datastore.config.AutoCreateTables {
		columnsList := keys(fixture.Records[0])
		if err := datastore.createTable(fixture.Table, columnsList); err != nil {
			log.Fatal(err)
		}
	}
	for _, record := range fixture.Records {
		query := buildNamedQuery(fixture.Table, record)
		fmt.Println(query)
	}
	return nil
}

func (datastore *postgresDatastore) Close() {
}

func buildConnectionString(conf entity.DBConfig) string {
	return fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		conf.Host, conf.Database, conf.User, conf.Password)
}

func buildNamedQuery(table string, record map[string]interface{}) string {
	cols := getColumns(record)
	colsStr := ""
	valuesStr := ""
	for i, c := range cols {
		colsStr += c
		valuesStr += ":" + c
		if i != len(cols)-1 {
			colsStr += ","
			valuesStr += ","
		}
	}
	q := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)", table, colsStr, valuesStr)
	fmt.Println(q)
	return q
}

func getColumns(record map[string]interface{}) []string {
	cols := []string{}
	for key := range record {
		cols = append(cols, key)
	}
	return cols
}

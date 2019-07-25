package postgres

import (
	"fmt"
	"os"
	"testing"

	"github.com/schehata/gofixtures/v3/entity"
)


type product struct {
	Name string
	Slug string
}

func prepareTestData(numberOfRecords int) entity.Fixture {
	records := make([]map[string]interface{}, numberOfRecords)
	for i := 0; i < numberOfRecords; i++ {
		records[i] = map[string]interface{}{
			"name": fmt.Sprintf("Product %d", i),
			"slug": fmt.Sprintf("product_%d", i),
		}
	}
	fixture := entity.Fixture{
		Table:   "products",
		Records: records,
	}

	return fixture
}

func TestInsertion(t *testing.T) {
	const numOfRecords = 100
	fixture := prepareTestData(numOfRecords)

	dbConfig := getDBConfig()
	dataStore := postgresDatastore{
		config: dbConfig,
	}
	err := dataStore.Connect()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	// truncate table first
	err = dataStore.Clear()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	err = dataStore.Insert(fixture)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	var products []product
	err = dataStore.db.Select(&products,"SELECT * FROM products")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if len(products) != numOfRecords {
		t.Fatalf("Number of records in the table doesn't much fixture records, found %d, expected %d",
			len(products),
			numOfRecords)
	}

}

func BenchmarkInsertion(b *testing.B) {

	dbConfig := entity.DBConfig{
		Driver:   "postgres",
		Database: os.Getenv("GOFIXTURES_TEST_DB_NAME"),
		User:     os.Getenv("GOFIXTURES_TEST_DB_USER"),
		Password: os.Getenv("GOFIXTURES_TEST_DB_PASSWORD"),
		Host:     os.Getenv("GOFIXTURES_TEST_DB_HOST"),
	}
	datastore := New(dbConfig)
	err := datastore.Connect()
	if err != nil {
		b.Error(err)
		b.Fail()
	}
	for i := 0; i < b.N; i++ {
		fixture := prepareTestData(b.N)
		err := datastore.Insert(fixture)
		if err != nil {
			b.Error(err)
			b.Fail()
		}
	}
}

func TestClear(t *testing.T) {
	fixture := prepareTestData(100)

	dbConfig := getDBConfig()
	datastore := postgresDatastore{
		config: dbConfig,
	}
	err := datastore.Connect()
	if err != nil {
		t.Fatal(err)
	}
	err = datastore.Insert(fixture)
	if err != nil {
		t.Fatal(err)
	}

	datastore.Clear()
	var products []product
	err = datastore.db.Select(&products, "SELECT * FROM products")
	if err != nil {
		t.Error(err.Error())
		t.Fatal(err)
	}
	if len(products) != 0 {
		t.Fatal("Clear should delete all rows from tables")
	}
}

func getDBConfig() entity.DBConfig {
	dbConfig := entity.DBConfig{
		Driver:           "postgres",
		Database:         os.Getenv("GOFIXTURES_TEST_DB_NAME"),
		User:             os.Getenv("GOFIXTURES_TEST_DB_USER"),
		Password:         os.Getenv("GOFIXTURES_TEST_DB_PASSWORD"),
		Host:             os.Getenv("GOFIXTURES_TEST_DB_HOST"),
		AutoCreateTables: true,
	}
	return dbConfig
}

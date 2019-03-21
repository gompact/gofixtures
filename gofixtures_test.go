package gofixtures

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/ishehata/gofixtures/v3/entity"
)

func getConfig() entity.Config {
	dbConfig := entity.DBConfig{
		Driver:           "postgres",
		Database:         os.Getenv("GOFIXTURES_TEST_DB_NAME"),
		User:             os.Getenv("GOFIXTURES_TEST_DB_USER"),
		Password:         os.Getenv("GOFIXTURES_TEST_DB_PASSWORD"),
		Host:             os.Getenv("GOFIXTURES_TEST_DB_HOST"),
		AutoCreateTables: true,
	}
	return entity.Config{
		DB: dbConfig,
	}
}

type FakeModel struct {
	Name string
	Age  int
}

func TestGoFixtures(t *testing.T) {
	config := getConfig()
	gf, err := New(config)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fakeModel := FakeModel{"John Doe", 27}
	data, err := json.Marshal(fakeModel)
	buf := bytes.NewReader(data)
	input := entity.Input{
		Filename: "test.json",
		Type:     ".json",
		Data:     buf,
	}
	err = gf.Load([]entity.Input{input})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	// TODO: check that db has the data

	err = gf.Clear()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	// TODO: check table has 0 rows
}

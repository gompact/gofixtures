package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ishehata/gofixtures/dal"
	"github.com/ishehata/gofixtures/dal/postgres"
	"github.com/ishehata/gofixtures/feed/cli"
	"github.com/ishehata/gofixtures/parser"
	"github.com/ishehata/gofixtures/parser/csv"
	"github.com/ishehata/gofixtures/parser/json"
	"github.com/ishehata/gofixtures/parser/yaml"
)

var queries []string

const version = "2.0.0"

func getParser(forType string) (parser.Parser, error) {
	switch forType {
	case ".json":
		return json.New(), nil
	case ".yaml":
		return yaml.New(), nil
	case ".csv":
		return csv.New(), nil
	default:
		return nil, errors.New("unsupported input type, supported types are YAML, CSV and JSON")
	}
}

func main() {
	cmdArgs := os.Args
	if cmdArgs[1] == "version" {
		log.Printf("version: %s", version)
		return
	}
	// read input using CLI
	feeder := cli.New()
	confInput, err := feeder.ReadConfig()
	if err != nil {
		feeder.Error(err, true)
	}

	confParser, err := getParser(confInput.Type)
	if err != nil {
		feeder.Error(errors.New("failed to parse database configuration"), false)
		feeder.Error(err, true)
	}
	conf, err := confParser.ParseConfig(confInput.Data)
	if err != nil {
		feeder.Error(err, true)
	}

	// connect to database
	var datastore dal.Datastore
	switch conf.DB.Driver {
	case "postgres":
		datastore = postgres.New(conf.DB)
	default:
		feeder.Error(errors.New("unsupported database driver"), true)
	}
	feeder.Debug("attempting to connect to datastore...")
	err = datastore.Connect()
	if err != nil {
		feeder.Error(errors.New("failed to connection to datastore"), false)
		feeder.Error(err, true)
	}
	defer datastore.Close()
	feeder.Info("Connection to datastore has been established")

	// get the data that needs to be parsed
	feeder.Info("loading fixture files...")
	input, err := feeder.GetInput()
	if err != nil {
		feeder.Error(err, true)
	}

	// based on type of the data, determine which parser is going to be used
	numberOfInputs := 0
	successfulInputs := 0
	for _, i := range input {
		numberOfInputs++
		p, err := getParser(i.Type)
		if err != nil {
			feeder.Error(err, true)
		}
		// parse the input
		fixture, err := p.Parse(i.Data)
		if err != nil {
			feeder.Warning("Failed to parse input, proceeding to next input")
			feeder.Error(err, false)
			continue
		}
		// TODO: maybe find a better approach to pass the filename to all
		// parsers and they can use/or not the filename.

		// Special case for the csv, set the table name from the file name
		if i.Type == ".csv" {
			fixture.Table = i.Filename
		}
		err = datastore.Insert(fixture)
		if err != nil {
			feeder.Error(errors.New("Failed to insert to datastore, "), false)
			feeder.Error(err, false)
			continue
		}
		successfulInputs++
	}

	feeder.Success(fmt.Sprintf("Successfully inserted %d out of %d\n", successfulInputs, numberOfInputs))
}

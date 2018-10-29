package csv

import (
	"errors"
	"io"

	"encoding/csv"

	"github.com/ishehata/gofixtures/entity"
)

// New creates a new instance of CSV parser
func New() *CSVParser {
	return &CSVParser{}
}

// CSVParser parses csv files and return a DBConf or Fixture data
type CSVParser struct {
}

// ParseDBConf parses db configurations from a JSON input
func (parser *CSVParser) ParseDBConf(input io.Reader) (entity.DBConfig, error) {
	return entity.DBConfig{}, errors.New("Reading Database configurations from a csv file is not supported")
}

// Parse parses list of items written in JSON
func (parser *CSVParser) Parse(input io.Reader) (entity.Fixture, error) {
	fixture := entity.Fixture{}
	r := csv.NewReader(input)
	records, err := r.ReadAll()
	if err != nil {
		return fixture, err
	}
	// get list of column names
	headers := records[0]
	// loop over each record and create a map from it, with column names as
	// map keys
	for _, record := range records[1:] {
		m := make(map[string]interface{})
		for i, h := range headers {
			m[h] = record[i]
		}
		// append the record to the fixture entity
		fixture.Records = append(fixture.Records, m)
	}
	return fixture, nil
}

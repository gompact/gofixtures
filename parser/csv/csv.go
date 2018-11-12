package csv

import (
	"io"

	"encoding/csv"

	"github.com/ishehata/gofixtures/v3/entity"
)

// New creates a new instance of CSV parser
func New(config entity.CSVConfig) *Parser {
	return &Parser{
		config: config,
	}
}

// Parser parses csv files and return a DBConf or Fixture data
type Parser struct {
	config entity.CSVConfig
}

// Parse parses list of items written in JSON
func (parser *Parser) Parse(input io.Reader) (entity.Fixture, error) {
	fixture := entity.Fixture{}
	r := csv.NewReader(input)
	// check that its not an empty rune
	if parser.config.Delimiter != rune(-1) {
		r.Comma = parser.config.Delimiter
	}
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

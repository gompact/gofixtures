package json

import (
	"encoding/json"
	"io"

	"github.com/ishehata/gofixtures/entity"
)

// New creates a new instance of JSON parser
func New() *Parser {
	return &Parser{}
}

type Parser struct {
}

// ParseConfig parses db configurations from a JSON input
func (parser *Parser) ParseConfig(input io.Reader) (entity.Config, error) {
	var data entity.Config
	err := json.NewDecoder(input).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// Parse parses list of items written in JSON
func (parser *Parser) Parse(input io.Reader) (entity.Fixture, error) {
	fixture := entity.Fixture{}
	err := json.NewDecoder(input).Decode(&fixture)
	if err != nil {
		return fixture, err
	}
	return fixture, nil
}

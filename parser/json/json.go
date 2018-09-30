package json

import (
	"io"

	j "encoding/json"

	"github.com/ishehata/gofixtures/entity"
	"github.com/ishehata/gofixtures/parser"
)

// New creates a new instance of JSON parser
func New() parser.Parser {
	return &json{}
}

type json struct {
}

// ParseDBConf parses db configurations from a JSON input
func (parser *json) ParseDBConf(input io.Reader) (entity.DBConfig, error) {
	var data entity.DBConfig
	err := j.NewDecoder(input).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// Parse parses list of items written in JSON
func (parser *json) Parse(input io.Reader) (entity.Fixture, error) {
	fixture := entity.Fixture{}
	err := j.NewDecoder(input).Decode(&fixture)
	if err != nil {
		return fixture, err
	}
	return fixture, nil
}

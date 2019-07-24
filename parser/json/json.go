package json

import (
	"encoding/json"
	"io"

	"github.com/schehata/gofixtures/v3/entity"
)

// New creates a new instance of JSON parser
func New() *Parser {
	return &Parser{}
}

type Parser struct {
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

package yaml

import (
	"io"

	"github.com/ishehata/gofixtures/entity"
	y "gopkg.in/yaml.v2"
)

// New creates a new instance of YAML parser
func New() *Parser {
	return &Parser{}
}

type Parser struct {
}

// ParseConfig parses db configurations from a YAML input
func (parser *Parser) ParseConfig(input io.Reader) (entity.Config, error) {
	var data entity.Config
	err := y.NewDecoder(input).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// Parse parses list of items written in YAML
func (parser *Parser) Parse(input io.Reader) (entity.Fixture, error) {
	fixture := entity.Fixture{}

	err := y.NewDecoder(input).Decode(&fixture)
	if err != nil {
		return fixture, err
	}
	return fixture, nil
}

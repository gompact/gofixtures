package yaml

import (
	"io"

	"github.com/ishehata/gofixtures/v3/entity"
	yaml "gopkg.in/yaml.v2"
)

// New creates a new instance of YAML parser
func New() *Parser {
	return &Parser{}
}

type Parser struct {
}

// Parse parses list of items written in YAML
func (parser *Parser) Parse(input io.Reader) (entity.Fixture, error) {
	fixture := entity.Fixture{}

	err := yaml.NewDecoder(input).Decode(&fixture)
	if err != nil {
		return fixture, err
	}
	return fixture, nil
}

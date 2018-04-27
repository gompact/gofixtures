package yaml

import (
	"io"

	"github.com/emostafa/gofixtures/entity"
	"github.com/emostafa/gofixtures/parser"
	y "gopkg.in/yaml.v2"
)

// New creates a new instance of YAML parser
func New() parser.Parser {
	return &yaml{}
}

type yaml struct {
}

// ParseDBConf parses db configurations from a YAML input
func (parser *yaml) ParseDBConf(input io.Reader) (entity.DBConfig, error) {
	var data entity.DBConfig
	err := y.NewDecoder(input).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// Parse parses list of items written in YAML
func (parser *yaml) Parse(input io.Reader) (entity.Fixture, error) {
	fixture := entity.Fixture{}

	err := y.NewDecoder(input).Decode(&fixture)
	if err != nil {
		return fixture, err
	}
	return fixture, nil
}

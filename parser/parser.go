package parser

import (
	"errors"
	"io"

	"github.com/ishehata/gofixtures/v3/entity"
	"github.com/ishehata/gofixtures/v3/parser/csv"
	"github.com/ishehata/gofixtures/v3/parser/json"
	"github.com/ishehata/gofixtures/v3/parser/yaml"
)

// Parser interface defines the methods needs to be implemented by parsers
// to integrate with the app
type Parser interface {
	ParseConfig(io.Reader) (entity.Config, error)
	Parse(io.Reader) (entity.Fixture, error)
}

// New creates a new parser for the passed file type, if the file type
// is not supported an error will be returned
func New(fileType string, config entity.Config) (Parser, error) {
	switch fileType {
	case ".json":
		return json.New(), nil
	case ".yaml":
		return yaml.New(), nil
	case ".csv":
		return csv.New(config.CSV), nil
	default:
		return nil, errors.New("unsupported input type, supported types are YAML, CSV and JSON")
	}
}

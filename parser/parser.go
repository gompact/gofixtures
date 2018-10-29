package parser

import (
	"io"

	"github.com/ishehata/gofixtures/v3/entity"
)

// Parser interface defines the methods needs to be implemented by parsers
// to integrate with the app
type Parser interface {
	ParseConfig(io.Reader) (entity.Config, error)
	Parse(io.Reader) (entity.Fixture, error)
}

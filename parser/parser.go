package parser

import (
	"io"

	"github.com/ishehata/gofixtures/entity"
)

// Parser interface defines the methods needs to be implemented by parsers
// to integrate with the app
type Parser interface {
	ParseDBConf(io.Reader) (entity.DBConfig, error)
	Parse(io.Reader) (entity.Fixture, error)
}

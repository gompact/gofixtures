package gofixtures

import (
	"errors"

	"github.com/schehata/gofixtures/v3/dal"
	"github.com/schehata/gofixtures/v3/dal/fake"
	"github.com/schehata/gofixtures/v3/dal/postgres"
	"github.com/schehata/gofixtures/v3/entity"
	"github.com/schehata/gofixtures/v3/feed/file"
	"github.com/schehata/gofixtures/v3/logger"
	"github.com/schehata/gofixtures/v3/parser"
)


const VERSION = "3.0.0"

// GoFixtures struct holds the configuration and the datastore collection
// also includes the methods needed to connect to perform operations
// with the fixtures
type GoFixtures struct {
	Config    entity.Config
	datastore dal.Datastore
}

// New creates a new instance of GoFixtures and connect to datastore
// based on the passed configuration
func New(config entity.Config) (*GoFixtures, error) {
	// connect to database
	var datastore dal.Datastore
	switch config.DB.Driver {
	case "postgres":
		datastore = postgres.New(config.DB)
	case "fake":
		datastore = fake.New(config.DB)
	default:
		logger.Error("unsupported database driver")
		return nil, errors.New("unsupported database driver")
	}
	logger.Debug("attempting to connect to datastore...")
	err := datastore.Connect()
	if err != nil {
		logger.Error("failed to connection to datastore")
		logger.Error(err.Error())
		return nil, err
	}

	return &GoFixtures{
		Config:    config,
		datastore: datastore,
	}, nil
}

// Load inserts fixtures into the database after deserializing the data
func (lib *GoFixtures) Load(inputs []entity.Input) error {
	for _, input := range inputs {
		go func(lib *GoFixtures, input entity.Input) {
			p, err := parser.New(input.Type, lib.Config)
			if err != nil {
				logger.Error(err.Error())
				return
			}
			fixture, err := p.Parse(input.Data)
			if err != nil {
				logger.Error(err.Error())
				return
			}
			// TODO: maybe find a better approach to pass the filename to all
			// parsers and they can use/or not the filename.

			// Special case for the csv, set the table name from the file name
			if input.Type == ".csv" {
				fixture.Table = input.Filename
			}
			err = lib.datastore.Insert(fixture)
			if err != nil {
				logger.Error(err.Error())
				return
			}
		}(lib, input)
	}
	return nil
}

// LoadFromFiles parses a list of files and insert their to given datastore
func (lib *GoFixtures) LoadFromFiles(files []string) error {
	feeder := file.New()
	inputs, err := feeder.Read(files)
	if err != nil {
		return err
	}
	return lib.Load(inputs)
}

// Clear clears all the database tables
func (lib *GoFixtures) Clear() error {
	return lib.datastore.Clear()
}

// Version returns the current version of gofixtures
func (lib *GoFixtures) Version() string {
	return VERSION
}

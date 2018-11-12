package main

import (
	"os"

	"github.com/ishehata/gofixtures/entity"
	"github.com/ishehata/gofixtures/logger"
	yaml "gopkg.in/yaml.v2"
)

// ReadConfig reades the configurations file. Make sure
// to call ReadCommandLineFlags() first.
func ReadConfig(file string) (entity.Config, error) {

	config := entity.Config{}
	logger.Info("reading configuration file...")
	f, err := os.Open(file)
	if err != nil {
		return config, err
	}
	logger.Success("configuration file has been loaded successfully")

	// parse yaml
	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

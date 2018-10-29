package entity

import "io"

// Config resembles the structure of the configuration file
type Config struct {
	DB  DBConfig  `json:"db" yaml:"db"`
	CSV CSVConfig `json:"csv" yaml:"csv"`
}

// ConfigInput resembles data read from the configuration file (or any other input)
type ConfigInput struct {
	Type string // JSON or YAML
	Data io.Reader
}

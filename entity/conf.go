package entity

import "io"

// Config resembles the structure of the configuration file
type Config struct {
	DB  DBConfig  `json:"db" yaml:"db"`
	CSV CSVConfig `json:"csv" yaml:"csv"`
}

// DBConfig resembles the configurations needed to connect to a datastore
type DBConfig struct {
	Driver           string `json:"driver" yaml:"driver"`
	Database         string `json:"database" yaml:"database"`
	User             string `json:"user" yaml:"user"`
	Password         string `json:"password" yaml:"password"`
	Host             string `json:"host" yaml:"host"`
	Port             int    `json:"port" yaml:"port"`
	AutoCreateTables bool   `json:"auto_create_tables" yaml:"auto_create_tables"`
}

// ConfigInput resembles data read from the configuration file (or any other input)
type ConfigInput struct {
	Type string // JSON or YAML
	Data io.Reader
}

// CSVConfig contains structure of configurations made only for CSV parsing
type CSVConfig struct {
	Separator string `json:"separator" yaml:"separator"`
}

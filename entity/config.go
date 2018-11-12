package entity

// Config resembles the structure of the configuration file
type Config struct {
	DB  DBConfig  `json:"db" yaml:"db"`
	CSV CSVConfig `json:"csv" yaml:"csv"`
}

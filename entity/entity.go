package entity

import "io"

// Input Defines the attributes of single Input coming through feeders
type Input struct {
	Type     string    // e.g: json, yaml
	Filename string    // e.g: countries.json
	Data     io.Reader // reader that holds the fixtures
}

// Fixture defines how the input data (Input.Data) should look like
// as it going to parsed by one the parsers (YAML, JSON...etc)
// note: it represents a signle YAML or JSON file
type Fixture struct {
	Table   string                   `json:"table" yaml:"table"`
	Records []map[string]interface{} `json:"records" yaml:"records"`
}

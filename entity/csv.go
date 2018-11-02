package entity

// CSVConfig contains structure of configurations made only for CSV parsing
type CSVConfig struct {
	Delimiter rune `json:"delimiter" yaml:"delimiter"`
}

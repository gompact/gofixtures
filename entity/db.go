package entity

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

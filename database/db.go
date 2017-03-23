package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//DBConfig is a struct resembles the configurations needed to connected to database
type DBConfig struct {
	Database string
	User     string
	Password string
	Host     string
	Port     int
}

//ConnectDatabase connects to postgresql db based on parameters of DBConfig object
func ConnectDatabase(conf string) (*sql.DB, error) {
	db, err := sql.Open("postgres", conf)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

package database

import (
	"database/sql"

	"log"

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
func ConnectDatabase(conf string) *sql.DB {
	db, err := sql.Open("postgres", conf)
	if err != nil {
		log.Print(err)
		log.Fatal("Error: The data source arguments are not valid")
	}
	err = db.Ping()
	if err != nil {
		log.Print(err)
		log.Fatal("Error: Couldn't stablish connection with the database")
	}
	return db
}

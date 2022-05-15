package database

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Variables from our environment variables.
var (
	host     = os.Getenv("HOST")
	user     = os.Getenv("USER")
	password = os.Getenv("PASSWORD")
	dbName   = os.Getenv("NAME")
	dbPort   = os.Getenv("DB_PORT")
	dbURI    = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, dbPort, user, password, dbName)
)

// Connect Creates a connection to the db which returns the (db instance or an error).
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	//import postgres lib to this package
	_ "github.com/lib/pq"
)

//ConfigTest used as configuration for testing
type ConfigTest struct {
	DBUser     string `json:"DB_USER"`
	DBPassword string `json:"DB_PASSWORD"`
	DBHost     string `json:"DB_HOST"`
	DBPort     int    `json:"DB_PORT"`
	DBName     string `json:"DB_NAME"`
}

//database used as interface to connect
var database *sqlx.DB

//New used for setup connection to postgres
func New(dbname, user, password, host string, port int) error {
	// create a normal database connection through database/sql
	dsn := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable", dbname, user, password, host, port)
	db, _ := sqlx.Open("postgres", dsn)

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(10)

	err := db.Ping()
	if err != nil {
		return err
	}

	database = db

	return nil
}

//GetDB used for get current database connection
func GetDB() *sqlx.DB {
	return database
}

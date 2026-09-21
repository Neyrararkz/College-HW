package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := "host=localhost port=5432 user=postgres password=1710 dbname=notes sslmode=disable";

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error opening database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("error connecting to database:", err)
	}

	log.Println("PostgreSQL connected successfully")
}
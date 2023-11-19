package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
)

const connStr = "user=%s password=%s host=%s port=%d dbname=%s sslmode=disable"

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", fmt.Sprintf(connStr, user, password, host, port, dbname))
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the database")
}

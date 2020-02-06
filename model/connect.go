package model

import (
	"database/sql"
	"log"
	"os"
)

var connect *sql.DB

func Connection() *sql.DB {
	dbType := os.Getenv("DB_PROVIDER")
	dbUrl := os.Getenv("DB_URL")
	db, err := sql.Open(dbType, dbUrl)
	if err != nil {
		log.Fatal(err.Error)
		return nil
	}
	connect = db
	return db
}

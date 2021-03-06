package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var connect *sql.DB

// Connection to database
func Connection() *sql.DB {
	dbType := os.Getenv("DB_PROVIDER")
	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open(dbType, dbURL)
	if err != nil {
		fmt.Println("error panic!!!")
		log.Fatal(err)
		return nil
	}
	connect = db
	return db
}

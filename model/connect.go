package model

import (
	"database/sql"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> 536ab95e47415021ee7892fa7c0214ae1e5552a9
	"log"
	"os"
)

var connect *sql.DB

func Connection() *sql.DB {
	dbType := os.Getenv("DB_PROVIDER")
	dbUrl := os.Getenv("DB_URL")
<<<<<<< HEAD
	db, err := sql.Open(dbType, dbUrl+"/todo")
	if err != nil {
		fmt.Println("error panic!!!")
=======
	db, err := sql.Open(dbType, dbUrl)
	if err != nil {
>>>>>>> 536ab95e47415021ee7892fa7c0214ae1e5552a9
		log.Fatal(err.Error)
		return nil
	}
	connect = db
	return db
}

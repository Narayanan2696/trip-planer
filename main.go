package main

import (
	"fmt"
	"log"
	"net/http"
	"trip-planer/controller"
	"trip-planer/initializers"
	"trip-planer/model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("=======Trip Planer=======")
	if initializers.InitiateEnv() == true {
		router := controller.Register()
		db := model.Connection()
		err := db.Ping()
		if err != nil {
			fmt.Println("connection is not active")
			log.Fatal(err.Error)
		}
		defer db.Close() // defer is used to execute the statement end of the scope here last line of main()
		http.ListenAndServe("localhost:3000", router)
	} else {
		log.Fatal("error in loading env")
	}

}

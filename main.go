package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	controller "trip-planer/controller/api/v1"
	"trip-planer/initializers"
	"trip-planer/model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("=======Trip Planer=======")
	if initializers.InitiateEnv() == true {
		fmt.Println("Did not reach returned false")
		router := controller.Register()
		db := model.Connection()
		err := db.Ping()
		if err != nil {
			fmt.Println("connection is not active")
			log.Fatal(err)
		}
		defer db.Close() // defer is used to execute the statement end of the scope here last line of main()
		host := os.Getenv("DOMAIN_HOST")
		http.ListenAndServe(host, router)
	} else {
		log.Fatal("error in loading env")
	}

}

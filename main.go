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
		router := controller.Register()
		db := model.Connection()
		err := db.Ping()
		if err != nil {
			fmt.Println("connection is not active")
			log.Fatal(err)
		}
		defer db.Close() // defer is used to execute the statement end of the scope here last line of main()
		model.PrepareSchemas()
		fmt.Println(os.Getenv("DATABASE_URL"))
		fmt.Println("=======connected========")
		http.ListenAndServe(":"+os.Getenv("PORT"), router)
	} else {
		log.Fatal("error in loading env")
	}

}

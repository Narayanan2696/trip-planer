package main

import (
	"fmt"
	"log"
	"net/http"
	"trip-planer/controller"
	"trip-planer/initializers"
)

func main() {
	fmt.Println("=======Trip Planer=======")
	if initializers.InitiateEnv() == true {
		router := controller.Register()
		http.ListenAndServe("localhost:3000", router)
	} else {
		log.Fatal("error in loading env")
	}

}

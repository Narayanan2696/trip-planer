package controller

import (
	"github.com/gorilla/mux"
)

func Register() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/trip_details", PostTripDetails())
	router.HandleFunc("/files", PostFileName())
	return router
}

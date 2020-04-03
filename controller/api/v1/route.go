package controller

import (
	"trip-planer/controller"

	"github.com/gorilla/mux"
)

func Register() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/trip_details", controller.PostTripDetails())
	router.HandleFunc("/files", controller.PostFileName())
	router.HandleFunc("/api/v1/trip_details", PostTripDetails())
	return router
}

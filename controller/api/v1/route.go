package controller

import (
	"github.com/gorilla/mux"
)

func Register() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/trip_details", PostTripDetails())
	router.HandleFunc("/api/v1/places", GetPlaces())
	router.HandleFunc("/api/v1/fuel_prices", GetFuelPrices())
	return router
}

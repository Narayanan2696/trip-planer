package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"trip-planer/lib"
	"trip-planer/model"
	"trip-planer/service"
	"trip-planer/views"
)

func PostTripDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fmt.Println("Encountered POST method of PostTripDetails")
			data := views.TripDetailsRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			GeoCoordinates, err := lib.FetchGeocodes(data.Source, data.Destination)
			if err != nil {
				log.Fatal(err.Error)
			}
			distance := service.CalculateDistance(GeoCoordinates, data.Unit)
			milage := model.ReadMilage(data.Car, data.FuelType)
			fuel := service.FuelRequired(data.Unit, distance, milage)

			json.NewEncoder(w).Encode(views.TripDetailsResponse{math.Round(distance), data.Unit, fuel})
		}
	}
}

package controller

import (
	"encoding/json"
	"fmt"
	"log"
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
			distance := service.CalculateDistance(GeoCoordinates)
			milage := model.ReadMilage(data.Car)
			fuel := service.FuelRequired(distance, milage)

			json.NewEncoder(w).Encode(views.TripDetailsResponse{fuel, "gallons"})
		}
	}
}

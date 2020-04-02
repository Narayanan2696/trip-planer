package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"trip-planer/lib"
	"trip-planer/lib/render"
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
				log.Println(err.Error)
			}
			distance, err := service.CalculateDistance(GeoCoordinates, data.Unit)
			milage := model.ReadMilage(data.Car, data.FuelType)
			fuel := service.FuelRequired(data.Unit, distance, milage)

			render.JSON(w, err, views.TripDetailsResponse{math.Round(distance), data.Unit, fuel})
		}
	}
}

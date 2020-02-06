package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"trip-planer/lib"
	"trip-planer/views"
)

func PostTripDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fmt.Println("Encountered GET method of PostTripDetails")
			data := views.TripDetailsRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			GeoCoordinates, err := lib.FetchGeocodes(data.Source, data.Destination)
			if err != nil {
				log.Fatal(err.Error)
			}
			fmt.Println(GeoCoordinates)
			json.NewEncoder(w).Encode(data)
		}
	}
}

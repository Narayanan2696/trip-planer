package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"trip-planer/lib"
	"trip-planer/lib/caching"
	"trip-planer/lib/render"
	"trip-planer/service"
	"trip-planer/views"
	"trip-planer/views/external_apis"
)

/**
	* should be optimized and cleaned currently taking 3ms v2.0-beta.1
**/
func GetPlaces() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Println("Encountered GET method of Places")
			data := views.TripDetailsRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			var travelPath external_apis.TravelPath
			cachedData, found := cacheMemory.Get(data.Source + "_" + data.Destination + "_" + data.Car + "_" + data.FuelType)
			if found {
				chachedDetails = cachedData.(caching.TravelDetails) // type assertion
				locationDetails := getLocationDetails(data, chachedDetails)
				travelPath = lib.FetchPlaces(locationDetails[0], locationDetails[1])
				// milage := model.ReadMilage(data.Car, data.FuelType)
				places, err := service.FindPathStops(travelPath, data.Unit, data.Source, data.Destination)
				render.JSON(w, err, places)
			} else {
				GeoCoordinates, err := lib.FetchGeocodes(data.Source, data.Destination)
				if err != nil {
					log.Println(err.Error)
				}
				fmt.Println("places reached")
				// milage := model.ReadMilage(data.Car, data.FuelType)
				travelPath = lib.FetchPlaces(GeoCoordinates[0], GeoCoordinates[1])
				places, err := service.FindPathStops(travelPath, data.Unit, data.Source, data.Destination)
				render.JSON(w, err, places)
			}
		}
	}
}

func getLocationDetails(trip views.TripDetailsRequest, cached caching.TravelDetails) []views.LocationDetails {
	location := make([]views.LocationDetails, 0, 2)
	location = append(location, views.LocationDetails{trip.Source, cached.SourceCoordinates.Latitude, cached.SourceCoordinates.Longitude})
	location = append(location, views.LocationDetails{trip.Destination, cached.DestinationCoordinates.Latitude, cached.DestinationCoordinates.Longitude})
	return location
}

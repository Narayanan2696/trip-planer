package controller

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"trip-planer/lib"
	"trip-planer/lib/caching"
	"trip-planer/lib/render"
	"trip-planer/service"
	"trip-planer/views"
	"trip-planer/views/external_apis"
	"github.com/gorilla/schema"
)

/**
	* should be optimized and cleaned currently taking 3ms v2.0-beta.1
**/
func GetPlaces() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Println("Encountered GET method of Places")
			dataSchema := views.TripDetailsQueryRequest{}
			schema.NewDecoder().Decode(&dataSchema, r.URL.Query())
			data := views.TripDetailsRequest{dataSchema.Source, dataSchema.Destination, dataSchema.Car, dataSchema.FuelType, dataSchema.Unit}
			var travelPath external_apis.TravelPath
			cachedData, found := cacheMemory.Get(data.Source + "_" + data.Destination + "_" + data.Car + "_" + data.FuelType + data.Unit)
			var GeoCoordinates = make([]views.LocationDetails, 0, 2)
			if found {
				chachedDetails = cachedData.(caching.TravelDetails) // type assertion
				locationDetails := getLocationDetails(data, chachedDetails)
				travelPath = lib.FetchPlaces(locationDetails[0], locationDetails[1])
				// milage := model.ReadMilage(data.Car, data.FuelType)
				places, err := service.FindPathStops(travelPath, data.Unit, data.Source, data.Destination)
				render.JSON(w, err, places)
			} else {
				sourceGeo, err := lib.FetchGeocodes(data.Source)
				destinationGeo, err := lib.FetchGeocodes(data.Destination)
				GeoCoordinates = append(GeoCoordinates, views.LocationDetails{sourceGeo.Place, sourceGeo.Latitude, sourceGeo.Longitude})
				GeoCoordinates = append(GeoCoordinates, views.LocationDetails{destinationGeo.Place, destinationGeo.Latitude, destinationGeo.Longitude})
				if err != nil {
					log.Println(err)
				}
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

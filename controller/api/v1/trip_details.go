package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"trip-planer/lib"
	"trip-planer/lib/caching"
	"trip-planer/lib/render"
	"trip-planer/model"
	"trip-planer/service"
	"trip-planer/views"

	"github.com/patrickmn/go-cache"
)

var cacheMemory = caching.Cache()
var chachedDetails caching.TravelDetails

// PostTripDetails Create trip details
func PostTripDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fmt.Println("Encountered POST method of PostTripDetails")
			data := views.TripDetailsRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			cachedData, found := cacheMemory.Get(data.Source + "_" + data.Destination + "_" + data.Car + "_" + data.FuelType + "_" + data.Unit)
			var GeoCoordinates = make([]views.LocationDetails, 0, 2)
			if found {
				chachedDetails = cachedData.(caching.TravelDetails) // type assertion
				/**
					* to be cleaned
				**/
				fuels := make([]views.FuelClassification, 0, 2)
				fuels = append(fuels, views.FuelClassification{chachedDetails.Gallons, "gallons"})
				fuels = append(fuels, views.FuelClassification{chachedDetails.Liters, "liters"})

				render.JSON(w, nil, views.TripDetailsResponse{chachedDetails.Distance, data.Unit, fuels})
			} else {
				sourceGeo, err := lib.FetchGeocodes(data.Source)
				destinationGeo, err := lib.FetchGeocodes(data.Destination)
				GeoCoordinates = append(GeoCoordinates, views.LocationDetails{sourceGeo.Place, sourceGeo.Latitude, sourceGeo.Longitude})
				GeoCoordinates = append(GeoCoordinates, views.LocationDetails{destinationGeo.Place, destinationGeo.Latitude, destinationGeo.Longitude})
				if err != nil {
					log.Println(err)
				}
				distance, err := service.CalculateDistance(GeoCoordinates, data.Unit)
				milage := model.ReadMilage(data.Car, data.FuelType)
				fuel := service.FuelRequired(data.Unit, distance, milage)
				// configure cache
				key := data.Source + "_" + data.Destination + "_" + data.Car + "_" + data.FuelType + "_" + data.Unit
				sourceCoordinate := caching.Coordinate{GeoCoordinates[0].Latitude, GeoCoordinates[0].Longitude}
				destinationCoordinate := caching.Coordinate{GeoCoordinates[1].Latitude, GeoCoordinates[1].Longitude}
				chachedDetails = caching.TravelDetails{sourceCoordinate, destinationCoordinate, math.Round(distance), fuel[0].Quantity, fuel[1].Quantity}
				cacheMemory.Set(key, chachedDetails, cache.DefaultExpiration)

				render.JSON(w, err, views.TripDetailsResponse{math.Round(distance), data.Unit, fuel})
			}
		}
	}
}

package service

import (
	"math"
	"trip-planer/views"
)

const pi = 3.14159265 // pi = 22 / 7

func CalculateDistance(geocodes []views.LocationDetails) float64 {
	sourceLat := getRadians(geocodes[0].Latitude)
	sourceLng := getRadians(geocodes[0].Longitude)
	destinationLat := getRadians(geocodes[1].Latitude)
	destinationLng := getRadians(geocodes[1].Longitude)

	radius := 3963.0 // radius of earth in miles

	/**
		haversine formula
		to calculate the distance between two coordinates in miles
	**/
	deltaLat := destinationLat - sourceLat
	deltaLng := destinationLng - sourceLng

	value := (math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(sourceLat)*math.Cos(destinationLat)*(math.Pow(math.Sin(deltaLng/2), 2)))
	sqrtValue := 2 * math.Asin(math.Sqrt(value))

	return radius * sqrtValue
}

func getRadians(coordinate float64) float64 {
	/**
		below is the simplified form of
		(coordinate / 57.29577951)
		180 / pi = 57.29577951
	**/
	return (coordinate / (180 / pi))
}

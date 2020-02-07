package service

import (
	"fmt"
	"math"
	"trip-planer/views"
)

const pi = 3.14159265 // pi = 22 / 7

/**
	*following is the version 1 logic of finding distance between two coordinates (as of v1.1 master branch)
**/
// func CalculateDistance(geocodes []views.LocationDetails) float64 {
// 	sourceLat := getRadians(geocodes[0].Latitude)
// 	sourceLng := getRadians(geocodes[0].Longitude)
// 	destinationLat := getRadians(geocodes[1].Latitude)
// 	destinationLng := getRadians(geocodes[1].Longitude)

// 	fmt.Printf("Slat:%f\tSlng:%f\tDlat:%f\tDlng:%f\t", sourceLat, sourceLng, destinationLat, destinationLng)
// 	radius := 3963.0 // radius of earth in miles

// 	/**
// 		haversine formula
// 		to calculate the distance between two coordinates in miles
//		as of in https://www.geeksforgeeks.org/program-distance-two-points-earth/
// 	**/
// 	deltaLat := destinationLat - sourceLat
// 	deltaLng := destinationLng - sourceLng

// 	value := (math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(sourceLat)*math.Cos(destinationLat)*(math.Pow(math.Sin(deltaLng/2), 2)))
// 	sqrtValue := 2 * math.Asin(math.Sqrt(value))

// 	fmt.Printf("Distance:%f\n", radius*sqrtValue)
// 	return radius * sqrtValue
// }

func CalculateDistance(geocodes []views.LocationDetails) float64 {
	sourceLat := getRadians(geocodes[0].Latitude)
	sourceLng := getRadians(geocodes[0].Longitude)
	destinationLat := getRadians(geocodes[1].Latitude)
	destinationLng := getRadians(geocodes[1].Longitude)

	fmt.Printf("Slat:%f\tSlng:%f\tDlat:%f\tDlng:%f\t", sourceLat, sourceLng, destinationLat, destinationLng)
	radius := 3963.0 // radius of earth in miles

	/**
		haversine formula
		to calculate the distance between two coordinates in miles
		as of in https://www.movable-type.co.uk/scripts/latlong.html
	**/
	deltaLat := getRadians(destinationLat - sourceLat)
	lamdaLng := getRadians(destinationLng - sourceLng)

	value := (math.Pow(math.Sin(deltaLat/2), 2) + (math.Cos(sourceLat)*math.Cos(destinationLat))*math.Pow(math.Sin(lamdaLng/2), 2))
	sqrtValue := 2 * math.Atan2(math.Sqrt(value), math.Sqrt(1-value))

	fmt.Printf("Distance:%f\n", radius*sqrtValue)
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

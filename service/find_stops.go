package service

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"trip-planer/lib"
	"trip-planer/lib/errors"
	"trip-planer/views"
	"trip-planer/views/external_apis"
)

var UNIT = []string{"miles", "kilometers"}

func FindPathStops(paths external_apis.TravelPath, unit, source, destination string) ([]views.Places, error) {
	distanceIntervalEnv := os.Getenv("DISTANCE_INTERVAL_METER")
	distanceInterval, _ := strconv.ParseInt(distanceIntervalEnv, 10, 64)
	var distance views.DistanceEntity
	var duration views.DurationEntity
	places := make([]views.Places, 0, len(paths.Route.Steps)-1)
	if contains(UNIT, unit) {
		coveredDistance := paths.Route.Steps[0].Distance
		coveredDuration := paths.Route.Steps[0].Duration
		places = append(places, views.Places{views.DistanceEntity{0, "miles"}, views.DurationEntity{"00:00", "HH::MM"}, views.LocationDetails{source, paths.Route.Steps[0].StartPoint.Latitude, paths.Route.Steps[0].StartPoint.Longitude}})
		for i := 1; i < len(paths.Route.Steps); i++ {
			if (coveredDistance + paths.Route.Steps[i].Distance) > distanceInterval {
				distance, duration = calculateQuantity(coveredDistance, coveredDuration, unit)
				location := lib.ReverseGeocoding(paths.Route.Steps[i].StartPoint.Latitude, paths.Route.Steps[i].StartPoint.Longitude)
				places = append(places, views.Places{distance, duration, views.LocationDetails{location, paths.Route.Steps[i].StartPoint.Latitude, paths.Route.Steps[i].StartPoint.Longitude}})
				coveredDistance = paths.Route.Steps[i].Distance
				coveredDuration = paths.Route.Steps[i].Duration
			} else {
				coveredDistance += paths.Route.Steps[i].Distance
				coveredDuration += paths.Route.Steps[i].Duration
			}
		}
		places = append(places, views.Places{distance, duration, views.LocationDetails{destination, paths.Route.Steps[len(paths.Route.Steps)-1].EndPoint.Latitude, paths.Route.Steps[len(paths.Route.Steps)-1].EndPoint.Longitude}})
		fmt.Println(places)
		return places, nil
	} else {
		err := errors.New(errors.CustomError{404, "Not Found", unit + " is not a valid unit of distance"})
		return nil, err
	}
}

func calculateQuantity(distance, duration int64, unit string) (views.DistanceEntity, views.DurationEntity) {
	var distanceEnt views.DistanceEntity
	var durationEnt views.DurationEntity
	if unit == UNIT[0] {
		distanceEnt.Distance = float64(distance) / (1000 * 1.6) // calculation to covert meters to miles
		distanceEnt.Unit = "miles"
		durationEnt.Duration = GetDuration(duration)
		durationEnt.Unit = "HH:MM"
	} else {
		distanceEnt.Distance = float64(distance) / 1000 // calculation to covert meters to miles
		distanceEnt.Unit = "kilometers"
		durationEnt.Duration = GetDuration(duration)
		durationEnt.Unit = "HH:MM"
	}
	return distanceEnt, durationEnt
}

func GetDuration(duration int64) string {
	value := float64(duration) / 3600
	hours := math.Floor(value)
	remainder := math.Abs(hours - value)
	minutes := math.Floor(remainder * 60)
	return fmt.Sprint(hours) + ":" + fmt.Sprint(minutes)
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

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

/**
	* To clean and optimize logic and api calls v2.0-beta.1
**/
func FindPathStops(paths external_apis.TravelPath, unit, source, destination string) ([]views.Places, error) {
	distanceIntervalEnv := os.Getenv("DISTANCE_INTERVAL_METER")
	distanceInterval, _ := strconv.ParseInt(distanceIntervalEnv, 10, 64)
	var distance views.DistanceEntity
	var duration views.DurationEntity
	places := make([]views.Places, 0, len(paths.Route.Steps)-1)
	locationChannel := make(chan string)
	defer close(locationChannel)
	if contains(UNIT, unit) {
		coveredDistance := paths.Route.Steps[0].Distance
		totalDistance := paths.Route.Steps[0].Distance
		coveredDuration := paths.Route.Steps[0].Duration
		totalDuration := paths.Route.Steps[0].Duration
		places = append(places, views.Places{0.0, "00:00", views.DistanceEntity{0, unit}, views.DurationEntity{"00:00", "HH::MM"}, views.LocationDetails{source, paths.Route.Steps[0].StartPoint.Latitude, paths.Route.Steps[0].StartPoint.Longitude}})
		for i := 1; i < len(paths.Route.Steps); i++ {
			if (coveredDistance + paths.Route.Steps[i].Distance) > distanceInterval {
				distance, duration = calculateQuantity(coveredDistance, coveredDuration, unit)
				go lib.ReverseGeocoding(paths.Route.Steps[i].StartPoint.Latitude, paths.Route.Steps[i].StartPoint.Longitude, locationChannel)
				location := <-locationChannel // first control comes here since no data in channel it will execute ReverseGeocoding goroutine and then resumes
				places = append(places, views.Places{getDistance(totalDistance, unit), getDuration(totalDuration), distance, duration, views.LocationDetails{location, paths.Route.Steps[i].StartPoint.Latitude, paths.Route.Steps[i].StartPoint.Longitude}})
				coveredDistance, coveredDuration = paths.Route.Steps[i].Distance, paths.Route.Steps[i].Duration
			} else {
				coveredDistance += paths.Route.Steps[i].Distance
				coveredDuration += paths.Route.Steps[i].Duration
			}
			totalDistance += paths.Route.Steps[i].Distance
			totalDuration += paths.Route.Steps[i].Duration
		}
		finalDistance, finalDuration := calculateQuantity(coveredDistance, coveredDuration, unit)
		places = append(places, views.Places{getDistance(totalDistance, unit), getDuration(totalDuration), finalDistance, finalDuration, views.LocationDetails{destination, paths.Route.Steps[len(paths.Route.Steps)-1].EndPoint.Latitude, paths.Route.Steps[len(paths.Route.Steps)-1].EndPoint.Longitude}})
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
	distanceEnt.Distance, distanceEnt.Unit = getDistance(distance, unit), unit
	durationEnt.Duration, durationEnt.Unit = getDuration(duration), "HH:MM"
	return distanceEnt, durationEnt
}

func getDistance(distance int64, unit string) float64 {
	if unit == UNIT[0] {
		return math.Round((float64(distance)/(1000*1.6))*100) / 100 // calculation to covert meters to miles
	} else {
		return math.Round((float64(distance)/1000)*100) / 100 // calculation to covert meters to miles
	}
}
func getDuration(duration int64) string {
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

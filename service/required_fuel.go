package service

import (
	"math"
	"strings"
	"trip-planer/views"
)

func FuelRequired(distanceUnit string, distance, milage float64) []views.FuelClassification {
	thresholdDistance := distance + (distance * 0.1)
	newMilage := distanceUnitMilage(distanceUnit, milage)
	gallons, liters := calculateFuelRequired(newMilage, thresholdDistance)
	// to return array of struct use make

	/**
		* to clean it in later versions
	**/
	fuels := make([]views.FuelClassification, 0, 2)
	fuels = append(fuels, views.FuelClassification{gallons, "gallons"})
	fuels = append(fuels, views.FuelClassification{liters, "liters"})
	return fuels
}

func distanceUnitMilage(distanceUnit string, milage float64) float64 {
	unit := strings.ToLower(distanceUnit)
	/**
		* to clean it in later versions
	**/
	if unit == "kilometers" {
		return 1.61 * milage
	} else {
		return milage
	}
}

func calculateFuelRequired(milage float64, thresholdDistance float64) (float64, float64) {
	/**
		* to clean it in later versions
	**/
	gallons := math.Ceil(thresholdDistance / milage)
	liters := math.Ceil((thresholdDistance * 3.79) / milage)
	return gallons, liters // fuel required in gallons
}

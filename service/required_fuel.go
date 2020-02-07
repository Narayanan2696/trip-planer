package service

import "math"

func FuelRequired(distance, milage float64) float64 {
	thresholdDistance := distance + (distance * 0.1)
	return math.Ceil(thresholdDistance / milage) // fuel required in gallons
}

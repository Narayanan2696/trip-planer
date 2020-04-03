package caching

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func Cache() *cache.Cache {
	return cache.New(24*time.Hour, 24*time.Hour)
}

type TravelDetails struct {
	SourceCoordinates      Coordinate
	DestinationCoordinates Coordinate
	Distance               float64
	Gallons                float64
	Liters                 float64
}

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

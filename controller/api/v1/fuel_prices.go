package controller

import (
	"fmt"
	"net/http"
	"trip-planer/lib"
)

func GetFuelPrices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Println("Encountered GET method of fuel prices")
			lib.FetchFuelPrices()
		}
	}
}

package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"trip-planer/lib/errors"
	"trip-planer/views"
	"trip-planer/views/external_apis"
)

func FetchGeocodes(place string) (views.LocationDetails, error) {
	key := os.Getenv("LOCATIONIQ_API_KEY")
	endpoint := os.Getenv("GEOCODING_ENDPOINT")
	url := endpoint + "?key=" + key + "&q=" + place + "&format=json"
	req, _ := http.Get(url)
	body, _ := ioutil.ReadAll(req.Body)
	var location []external_apis.Geocode
	json.Unmarshal(body, &location)
	fmt.Println(location[0])
	var placeDetail views.LocationDetails
	if location[0].Lat == "" || location[0].Lon == "" {
		return placeDetail, errors.New(errors.CustomError{404, "NOT_FOUND", "Latitude:" + location[0].Lat + " or Longitude: " + location[0].Lon + " is empty"})
	} else {
		latitude, _ := strconv.ParseFloat(location[0].Lat, 64)
		longitude, _ := strconv.ParseFloat(location[0].Lon, 64)
		return views.LocationDetails{place, latitude, longitude}, nil
	}
}

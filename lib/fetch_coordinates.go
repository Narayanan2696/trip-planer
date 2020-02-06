package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"trip-planer/views"
)

func FetchGeocodes(source, destination string) ([]views.LocationDetails, error) {
	GeoCoordinates := []views.LocationDetails{}
	sourceCoordinates, err := getCoordinates(source)
	destinationCoordinates, err := getCoordinates(destination)
	GeoCoordinates = append(GeoCoordinates, sourceCoordinates, destinationCoordinates)
	if err != nil {
		return nil, err
	}
	return GeoCoordinates, nil
}

func getCoordinates(location string) (views.LocationDetails, error) {
	key := os.Getenv("GEOCODING_API_KEY")
	endpoint := os.Getenv("GEOCODING_ENDPOINT")
	latIndex, err := strconv.ParseInt(os.Getenv("LATITUDE_INDEX"), 10, 64)
	lngIndex, err := strconv.ParseInt(os.Getenv("LONGITUDE_INDEX"), 10, 64)
	url := endpoint + "?key=" + key + "&location=" + location
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error)
	}
	responseBody, _ := ioutil.ReadAll(response.Body)
	strResponse := string(responseBody)
	marshalledResponse := map[string]interface{}{}
	if err := json.Unmarshal([]byte(strResponse), &marshalledResponse); err != nil {
		panic(err)
	}
	for _, v := range marshalledResponse {
		if _, ok := v.(map[string]interface{}); ok {
		} else {
			parseString := "latLng:map[lat:]"
			str := fmt.Sprintf("%v", v)
			arr := strings.Fields(str)
			lat := strings.Trim(arr[latIndex], parseString)
			lng := strings.Trim(arr[lngIndex], parseString)
			latitude, err := strconv.ParseFloat(lat, 64)
			longitude, err := strconv.ParseFloat(lng, 64)
			if err != nil {
				return views.LocationDetails{}, err
			}
			return views.LocationDetails{location, latitude, longitude}, nil
		}
	}
	return views.LocationDetails{}, nil
}

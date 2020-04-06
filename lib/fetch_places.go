package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"trip-planer/views"
	"trip-planer/views/external_apis"
)

func FetchPlaces(source, destination views.LocationDetails) external_apis.TravelPath {
	key := os.Getenv("RAPID_API_KEY")
	endpoint := os.Getenv("PLACES_ENDPOINT")
	const domain = "trueway-directions2.p.rapidapi.com"
	url := endpoint + "?origin=" + fmt.Sprint(source.Latitude) + "%252C" + fmt.Sprint(source.Longitude) + "&destination=" + fmt.Sprint(destination.Latitude) + "%252C" + fmt.Sprint(destination.Longitude)
	fmt.Println(url)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", domain)
	req.Header.Add("x-rapidapi-key", key)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var travelDetails external_apis.TravelPath
	json.Unmarshal(body, &travelDetails)
	return travelDetails
}

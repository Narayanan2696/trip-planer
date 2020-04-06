package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"trip-planer/views/external_apis"
)

func ReverseGeocoding(latitude, longitude float64, channel chan string) {
	key := os.Getenv("LOCATIONIQ_API_KEY")
	endpoint := os.Getenv("REVERSE_GEOCODING_ENDPOINT")
	url := endpoint + "?key=" + key + "&lat=" + fmt.Sprint(latitude) + "&lon=" + fmt.Sprint(longitude) + "&format=json"
	fmt.Println(url)
	req, _ := http.Get(url)
	body, _ := ioutil.ReadAll(req.Body)
	var address external_apis.ReverseGeocode
	json.Unmarshal(body, &address)
	channel <- address.DisplayName
}

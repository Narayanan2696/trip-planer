package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func FetchFuelPrices() {
	// key := os.Getenv("RAPID_API_KEY")
	endpoint := os.Getenv("AMENITIES_ENDPOINT")
	url := endpoint + "?mylat=51.50784&amenity=fuel&mylon=-0.127324&mode=undefined"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "community-amenities-maps.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "fa0688c241mshdd93a3e1da8af2ep1b23d9jsne372c7cd11ad")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fileWrite(string(body))
}

func fileWrite(content string) {
	f, err := os.Create("/Users/narayananv/Desktop/go_data_files/amenities.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

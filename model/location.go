package model

import (
	"log"
	"trip-planer/views"
)

func ReadLocation(place string) (views.LocationDetails, error) {
	rows, err := connect.Query("SELECT PLACE, LATITUDE, LONGITUDE FROM LOCATION WHERE PLACE=?", place)
	defer rows.Close()
	location := views.LocationDetails{}
	if err != nil {
		log.Fatal(err)
		return location, err
	}
	for rows.Next() {
		rows.Scan(&location.Place, &location.Latitude, &location.Longitude)
	}
	return location, err
}

// InsertLocation create a location
func InsertLocation(location views.LocationDetails) {
	insertQ, err := connect.Query("INSERT INTO LOCATION(PLACE, LATITUDE, LONGITUDE) VALUES(?,?,?)", location.Place, location.Latitude, location.Longitude)
	defer insertQ.Close()
	if err != nil {
		log.Fatal(err)
		log.Fatal("location got through API call was not successfully loaded in DB!!")
	}
}

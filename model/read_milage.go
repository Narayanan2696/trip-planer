package model

import (
	"log"
	"strconv"
	"fmt"
)

func ReadMilage(car, fuelType string) float64 {
	rows, err := connect.Query(`SELECT MPG FROM AUTOMOBILEDETAILS WHERE FUEL_TYPE = $1 AND CAR_NAME = $2`, fuelType, car)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	milage := milage{}
	for rows.Next() {
		var mpg string
		rows.Scan(&mpg)
		floatMPG, _ := strconv.ParseFloat(mpg, 64)
		milage.mpg = floatMPG
	}
	fmt.Printf("miles per gallon:%f\n", milage.mpg)
	return milage.mpg
}

type milage struct {
	mpg float64
}

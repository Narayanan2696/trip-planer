package model

import (
	"log"
)

func ReadMilage(car, fuelType string) float64 {
	rows, err := connect.Query("SELECT MPG FROM FUELMILAGE WHERE TYPE = ? AND MILAGE_ID IN (SELECT ID FROM MILAGE WHERE CAR_NAME = ?)", fuelType, car)
	defer rows.Close()
	if err != nil {
		log.Fatal(err.Error)
	}
	milage := milage{}
	for rows.Next() {
		rows.Scan(&milage.mpg)
	}
	// fmt.Printf("miles per gallon:%f\n", milage.mpg)
	return milage.mpg
}

type milage struct {
	mpg float64
}

package model

import (
	"log"
)

func ReadMilage(car string) float64 {
	rows, err := connect.Query("SELECT MPG FROM MILAGE WHERE CAR_NAME=?", car)
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

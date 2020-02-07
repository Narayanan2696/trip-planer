package views

type TripDetailsRequest struct {
	Source      string `json:"from"`
	Destination string `json:"to"`
	Car         string `json:"car_name"`
}

type TripDetailsResponse struct {
	Fuel float64 `json:"fuel"`
	Unit string  `json:"unit"`
}

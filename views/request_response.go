package views

type TripDetailsRequest struct {
	Source      string `json:"from"`
	Destination string `json:"to"`
	Car         string `json:"car_name"`
}

type TripDetailsResponse struct {
	Fuel string `json:"fuel"`
}

package views

type TripDetailsRequest struct {
	Source      string `json:"from"`
	Destination string `json:"to"`
}

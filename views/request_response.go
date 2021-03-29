package views

type TripDetailsRequest struct {
	Source      string `json:"from"`
	Destination string `json:"to"`
	Car         string `json:"car_name"`
	FuelType    string `json:"fuel_type"`
	Unit        string `json:"unit"`
}

type TripDetailsQueryRequest struct {
	Source      string `schema:"from"`
	Destination string `schema:"to"`
	Car         string `schema:"car_name"`
	FuelType    string `schema:"fuel_type"`
	Unit        string `schema:"unit"`
}

type TripDetailsResponse struct {
	Distance float64              `json:"distance"`
	Unit     string               `json:"unit"`
	Fuel     []FuelClassification `json:"fuel"`
}

type FuelClassification struct {
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

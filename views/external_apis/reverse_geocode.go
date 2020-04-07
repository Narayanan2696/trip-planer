package external_apis

type ReverseGeocode struct {
	PlaceId     string   `json:"place_id"`
	Licence     string   `json:"licence"`
	OsmType     string   `json:"osm_type"`
	OsmID       string   `json:"osm_id"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	DisplayName string   `json:"display_name"`
	Address     address  `json:"address"`
	BoundingBox []string `json:"boundingbox"`
}

type address struct {
	Road           string `json:"road"`
	Village        string `json:"village"`
	VillageCountry string `json:"county"`
	District       string `json:"state_district"`
	State          string `json:"state"`
	PostCode       string `json:"postcode"`
	Country        string `json:"country"`
	CountryCode    string `json:"country_code"`
}

package external_apis

type TravelPath struct {
	Route routeDetails `json:"route"`
}

type routeDetails struct {
	Distance int64       `json:"distance"`
	Duration int64       `json:"duration"`
	Bounds   bound       `json:"bounds"`
	Geometry coordinates `json:"geometry"`
	Steps    []step      `json:"steps"`
}

type coordinates struct {
	Coordinates [][]float64 `json:"coordinates"`
}

type step struct {
	Distance        int64  `json:"distance"`
	Duration        int64  `json:"duration"`
	StartPointIndex int64  `json:"start_point_index"`
	StartPoint      latlng `json:"start_point"`
	EndPointIndex   int64  `json:"end_point_index"`
	EndPoint        latlng `json:"end_point"`
	Bounds          bound  `json:"bounds"`
}

type bound struct {
	South float64 `json:"south"`
	West  float64 `json:"west"`
	North float64 `json:"north"`
	East  float64 `json:"east"`
}

type latlng struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

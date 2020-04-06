package views

type Places struct {
	CumulativeDistance float64         `json:"cumulative_distance"`
	CumulativeDuration string          `json:"cumulative_duration"`
	Distance           DistanceEntity  `json:"distance"`
	Duration           DurationEntity  `json:"duration"`
	Location           LocationDetails `json:"location"`
}

type DistanceEntity struct {
	Distance float64 `json:"distance"`
	Unit     string  `json:"unit"`
}

type DurationEntity struct {
	Duration string `json:"duration"`
	Unit     string `json:"unit"`
}

package models

type ResponseListToStation struct {
	LinePlatform string  `json:"line_platform"`
	LineName     string  `json:"line_name"`
	StationCode  string  `json:"station_code"`
	StationName  string  `json:"station_name"`
	MinPrice     float64 `json:"station_min_price"`
	MinTime      int     `json:"station_min_time"`
	MinStation   int     `json:"station_min_station"`
}

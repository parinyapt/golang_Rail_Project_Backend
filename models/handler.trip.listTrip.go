package models

type ResponseListTrip struct {
	FromLinePlatform string  `json:"from_line_platform"`
	FromLineName     string  `json:"from_line_name"`
	FromStationCode  string  `json:"from_station_code"`
	FromStationName  string  `json:"from_station_name"`
	ToLinePlatform   string  `json:"to_line_platform"`
	ToLineName       string  `json:"to_line_name"`
	ToStationCode    string  `json:"to_station_code"`
	ToStationName    string  `json:"to_station_name"`
	MinPrice         float64 `json:"price"`
	MinTime          int     `json:"time"`
	MinStation       int     `json:"station"`
	TripName         string  `json:"trip_name"`
	TripID           int     `json:"trip_id"`
}

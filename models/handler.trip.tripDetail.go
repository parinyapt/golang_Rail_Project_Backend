package models

type ResponseTripDetail struct {
	FromLinePlatform string        `json:"from_line_platform"`
	FromLineName     string        `json:"from_line_name"`
	FromStationCode  string        `json:"from_station_code"`
	FromStationName  string        `json:"from_station_name"`
	ToLinePlatform   string        `json:"to_line_platform"`
	ToLineName       string        `json:"to_line_name"`
	ToStationCode    string        `json:"to_station_code"`
	ToStationName    string        `json:"to_station_name"`
	MinPrice         float64       `json:"price"`
	MinTime          int           `json:"time"`
	MinStation       int           `json:"station"`
	TripName         string        `json:"trip_name"`
	PlaceDetail      []PlaceDetail `json:"place_detail"`
}

type PlaceDetail struct {
	LineID         int    `json:"line_id"`
	LinePlatform   string `json:"line_platform"`
	LineColor      string `json:"line_color"`
	LineName       string `json:"line_name"`
	StationID      int    `json:"station_id"`
	StationCode    string `json:"station_code"`
	StationName    string `json:"station_name"`
	PlaceID        string `json:"place_id"`
	PlaceName      string `json:"place_name"`
	PlaceLatitude  string `json:"place_latitude"`
	PlaceLongitude string `json:"place_longitude"`
	PlaceDistance  string `json:"place_distance"`
	PlaceImage     string `json:"place_image"`
}

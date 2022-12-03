package models

type ResponseRouteDetailData struct {
	FromLinePlatform string                `json:"from_line_platform"`
	FromLineName     string                `json:"from_line_name"`
	FromStationCode  string                `json:"from_station_code"`
	FromStationName  string                `json:"from_station_name"`
	ToLinePlatform   string                `json:"to_line_platform"`
	ToLineName       string                `json:"to_line_name"`
	ToStationCode    string                `json:"to_station_code"`
	ToStationName    string                `json:"to_station_name"`
	MinPrice         float64               `json:"min_price"`
	MinTime          int                   `json:"min_time"`
	MinStation       int                   `json:"min_station"`
	StationList      []ResponseRouteDetail `json:"station_list"`
}

type ResponseRouteDetail struct {
	LinePlatform string `json:"line_platform"`
	StationCode  string `json:"station_code"`
	StationName  string `json:"station_name"`
}

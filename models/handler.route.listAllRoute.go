package models

type ResponseListAllRoute struct {
	FromLinePlatform string         `json:"from_line_platform"`
	FromLineName     string         `json:"from_line_name"`
	FromStationCode  string         `json:"from_station_code"`
	FromStationName  string         `json:"from_station_name"`
	ToLinePlatform   string         `json:"to_line_platform"`
	ToLineName       string         `json:"to_line_name"`
	ToStationCode    string         `json:"to_station_code"`
	ToStationName    string         `json:"to_station_name"`
	MinPrice         float64        `json:"min_price"`
	MinTime          int            `json:"min_time"`
	MinStation       int            `json:"min_station"`
	AllRouteList     []AllRouteList `json:"route"`
}

type AllRouteList struct {
	Platform []string `json:"platform"`
	Price    float64  `json:"price"`
	Time     int      `json:"time"`
	Station  int      `json:"station"`
}

type DBStructAllRouteList struct {
	Platform string `json:"platform"`
	Price    float64  `json:"price"`
	Time     int      `json:"time"`
	Station  int      `json:"station"`
}

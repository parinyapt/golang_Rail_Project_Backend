package models

type ResponseListStation struct {
	StationID        int    `json:"station_id"`
	StationCode      string `json:"station_code"`
	StationName      string `json:"station_name"`
	StationLatitude  string `json:"station_latitude"`
	StationLongitude string `json:"station_longitude"`
	StationGoogleMap string `json:"station_googlemap"`
	StationImage     string `json:"station_image"`
}

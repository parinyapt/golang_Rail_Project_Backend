package models

type ResponseStationDetail struct {
	StationID        int    `json:"station_id"`
	StationCode      string `json:"station_code"`
	StationName      string `json:"station_name"`
	StationLatitude  string `json:"station_latitude"`
	StationLongitude string `json:"station_longitude"`
	StationGoogleMap string `json:"station_googlemap"`
	StationImage     string `json:"station_image"`
	StationExit      []StationExitDetail
	StationFacility  []StationFacilityDetail
}

type StationExitDetail struct {
	ExitNumber string `json:"exit_number"`
	ExitName   string `json:"exit_name"`
}

type StationFacilityDetail struct {
	FacilityIcon string `json:"facility_icon"`
	FacilityName string `json:"facility_name"`
}
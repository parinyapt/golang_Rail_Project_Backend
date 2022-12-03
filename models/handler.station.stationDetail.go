package models

type ResponseStationDetail struct {
	StationID        int                     `json:"station_id"`
	StationCode      string                  `json:"station_code"`
	StationName      string                  `json:"station_name"`
	StationLatitude  string                  `json:"station_latitude"`
	StationLongitude string                  `json:"station_longitude"`
	StationGoogleMap string                  `json:"station_googlemap"`
	StationImage     string                  `json:"station_image"`
	StationExit      []StationExitDetail     `json:"station_exit"`
	StationFacility  []StationFacilityDetail `json:"station_facility"`
	StationLink      []StationLinkDetail     `json:"station_link"`
}

type StationExitDetail struct {
	ExitNumber string `json:"exit_number"`
	ExitName   string `json:"exit_name"`
}

type StationFacilityDetail struct {
	FacilityIcon string `json:"facility_icon"`
	FacilityName string `json:"facility_name"`
}

type StationLinkDetail struct {
	StationCode  string `json:"station_code"`
	StationName  string `json:"station_name"`
	LinePlatform string `json:"line_platform"`
	LineName     string `json:"line_name"`
}

// SELECT crp_station.station_code, ctx1.translation_text, crp_line.line_platform, ctx2.translation_text FROM `crp_station_link` INNER JOIN crp_station ON crp_station_link.link_station_id_to = crp_station.station_id INNER JOIN crp_line ON crp_station.station_line_id = crp_line.line_id INNER JOIN crp_translation ctx1 ON crp_station.station_name_tx_id = ctx1.translation_id INNER JOIN crp_translation ctx2 ON crp_line.line_name_tx_id = ctx2.translation_id INNER JOIN crp_language ON ctx1.translation_language_id = crp_language.language_id AND ctx2.translation_language_id = crp_language.language_id WHERE crp_language.language_code = ? AND crp_station_link.link_station_id_from = ?;
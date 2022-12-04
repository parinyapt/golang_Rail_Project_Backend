package models

type ResponseListPlace struct {
	PlaceID        string `json:"place_id"`
	PlaceName      string `json:"place_name"`
	PlaceLatitude  string `json:"place_latitude"`
	PlaceLongitude string `json:"place_longitude"`
	PlaceDistance  string `json:"place_distance"`
	PlaceImage     string `json:"place_image"`
}

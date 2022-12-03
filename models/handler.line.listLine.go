package models

type ResponseListLine struct {
	LineID int `json:"line_id"`
	LinePlatform string `json:"line_platform"`
	LineColor string `json:"line_color"`
	LineName string `json:"line_name"`
}
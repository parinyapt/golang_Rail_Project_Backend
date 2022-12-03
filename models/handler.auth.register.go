package models

type RequestRegister struct {
	RefID string `json:"ref_id" validate:"required,uuid"`
	Name  string `json:"name" validate:"required"`
}

package models

type RequestrequestOTP struct {
	Email string `json:"email" validate:"required,email"`
}

type ResponserequestOTP struct {
	RefID string `json:"ref_id"`
	Status string `json:"status"`
}
package models

type RequestLogin struct {
	Email   string `json:"email" validate:"required,email"`
	OTPCode string `json:"otp_code" validate:"required,numeric,len=6"`
	RefID   string `json:"ref_id" validate:"required,uuid"`
}

type ResponseLogin struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}

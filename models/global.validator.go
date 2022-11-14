package models

type ValidatorError struct {
	Field    string `json:"field"`
	ErrorMsg string `json:"error_message"`
}

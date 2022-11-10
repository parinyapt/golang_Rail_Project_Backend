package models

type ParameterSendMail struct {
	Mailto   []string
	Subject  string
	Body     string
	BodyType string
}

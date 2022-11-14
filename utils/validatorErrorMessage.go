package utils

func ValidatorErrorMsg(tag string) string {
	switch tag {
	case "required":
			return "This field is required"
	case "email":
			return "Invalid email"
	}
	return "Invalid Format"
}
package utils

import (
	"os"

	"github.com/pkg/errors"
	gomail "gopkg.in/gomail.v2"
	
	"github.com/parinyapt/Rail_Project_Backend/models"
)



func SendMail(config models.ParameterSendMail) error {
	if (config.BodyType != "html" && config.BodyType != "plain"){
		return errors.Wrap(errors.New("[Custom]->Body Type Not Exist"), "[Error]->Body Type Check")
	}

	m := gomail.NewMessage()
	m.SetHeader("From",os.Getenv("SENDMAIL_FROM_ALIAS_NAME") + " <" + os.Getenv("SENDMAIL_FROM_EMAIL") + ">")
	m.SetHeader("To", config.Mailto...)
	m.SetHeader("Subject", config.Subject)
	m.SetBody("text/" + config.BodyType, config.Body)

	d := gomail.NewDialer("smtp.gmail.com", 465, os.Getenv("SENDMAIL_USERNAME"), os.Getenv("SENDMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return errors.Wrap(err, "[Error]->Send Email Error")
	}
	return nil
}
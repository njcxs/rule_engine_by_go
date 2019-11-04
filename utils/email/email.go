package email

import (
	"github.com/go-gomail/gomail"
	"log"
)

type EmailConf struct {

	ServerHost string
	ServerPort int
	EmailAddr string
	Passwd string
}

type ToSomeBody struct {

	To []string
	Cc []string
}


func NewEmailSender(emailconf *EmailConf) *gomail.Dialer {

	d := gomail.NewPlainDialer(emailconf.ServerHost, emailconf.ServerPort, emailconf.EmailAddr, emailconf.Passwd)

	_ , err := d.Dial()
	if err != nil {
		log.Println("Could not send email,smtp configure is Incorrect.")
	}
	return d
}


func SendEmail(d *gomail.Dialer, to *ToSomeBody, Subject string, content string) error{

		m := gomail.NewMessage()
		m.SetHeader("From", cfg.Smtp.From)
		m.SetHeader("To", to.To...)
	    m.SetHeader("Cc", to.Cc...)
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", content)

		if err := d.DialAndSend(m); err != nil {

		}

		return err
}

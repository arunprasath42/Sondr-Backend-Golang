package notification

import (
	"crypto/tls"
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendMail(subject, body string, toEmails ...string) (string, error) {
	fmt.Println("Mail service Called")
	m := gomail.NewMessage()
	m.SetHeader("From", viper.GetString("sondr-admin.email"))
	m.SetHeader("To", toEmails...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)
	d := gomail.NewDialer(viper.GetString("sondr-admin.host"), 587, viper.GetString("sondr-admin.email"), viper.GetString("sondr-admin.password"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return "", err
	}
	return "Email Send Successfully", nil
}

func HtmlMail(subject, body string, toEmails ...string) (string, error) {
	fmt.Println("Mail service Called")
	m := gomail.NewMessage()
	m.SetHeader("From", viper.GetString("sondr-admin.email"))
	m.SetHeader("To", toEmails...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(viper.GetString("sondr-admin.host"), 587, viper.GetString("sondr-admin.email"), viper.GetString("sondr-admin.password"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return "", err
	}
	return "Email Send Successfully", nil
}

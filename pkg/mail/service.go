package mail

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendMail(to, subject, code string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./pkg/configs")

	server, port, username, password := viper.GetString("mail.server"), viper.GetInt("mail.smtpPort"), viper.GetString("mail.username"), viper.GetString("mail.password")

	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", `Hello Verification code is `+code+` `)

	d := gomail.NewDialer(server, port, username, password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

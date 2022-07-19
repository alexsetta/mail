package mail

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"strings"
)

type MailConfig struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

type Config struct {
	Configs []MailConfig `json:"configs"`
}

func ReadConfig(fileName string) (MailConfig, error) {
	var config Config
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return MailConfig{}, fmt.Errorf("readConfig: %w", err)
	}
	reader := strings.NewReader(string(b))

	if err := json.NewDecoder(reader).Decode(&config); err != nil {
		return MailConfig{}, fmt.Errorf("readConfig: %w", err)
	}
	c := config.Configs[0]
	return c, nil
}

func SendEmail(config MailConfig, emailTo, subject, msg string) error {
	if emailTo == "" || config.Login == "" || config.Password == "" {
		return nil
	}

	m := gomail.NewMessage()
	m.SetHeader("From", config.Login)
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", msg)

	d := gomail.NewDialer(config.Host, config.Port, config.Login, config.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("sendEmail: %w", err)
	}
	return nil
}

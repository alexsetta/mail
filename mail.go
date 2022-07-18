package mail

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"strings"
)

type Config struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func ReadConfig(fileName string) (Config, error) {
	cfg := Config{}
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return cfg, fmt.Errorf("readConfig: %w", err)
	}
	reader := strings.NewReader(string(b))

	if err := json.NewDecoder(reader).Decode(&cfg); err != nil {
		return cfg, fmt.Errorf("readConfig: %w", err)
	}
	return cfg, nil
}

func SendEmail(config Config, emailTo, subject, msg string) error {
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

package mail

import (
	"testing"
)

func TestReadConfig(t *testing.T) {
	_, err := ReadConfig("./config.json")
	if err != nil {
		t.Errorf("ReadConfig() error = %v", err)
		return
	}
}

func TestSendEmail(t *testing.T) {
	config, err := ReadConfig("./config.json")
	if err != nil {
		t.Errorf("ReadConfig() error = %v", err)
		return
	}

	if err := SendEmail(config, "alexsetta@gmail.com", "Teste de email", "Corpo do email de teste"); err != nil {
		t.Errorf("SendEmail() error = %v", err)
	}

}

package mail

import (
	"fmt"
	"testing"
)

func TestReadConfig(t *testing.T) {
	c, err := ReadConfig("./config.json")
	if err != nil {
		t.Errorf("ReadConfig() error = %v", err)
		return
	}
	fmt.Println(c)
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

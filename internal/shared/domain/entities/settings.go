package entities

import "fmt"

type Settings struct {
	Port        int
	PaymentsAPI string
}

func NewSettings(port int, paymentsAPI string) *Settings {
	return &Settings{port, paymentsAPI}
}

func (s Settings) GetPort() string {
	return fmt.Sprintf(":%d", s.Port)
}

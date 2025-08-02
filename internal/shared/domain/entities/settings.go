package entities

import "fmt"

type Settings struct {
	Port int
}

func NewSettings(port int) *Settings {
	return &Settings{port}
}

func (s Settings) GetPort() string {
	return fmt.Sprintf(":%d", s.Port)
}

package entities

import (
	"errors"
	"fmt"
	"time"
)

type Model struct {
	Name  string
	Year  int
	Brand Brand
}

func (m *Model) IsValid() error {
	if m.Name == "" {
		return errors.New("model name is required")
	}
	nextYear := time.Now().Year() + 1
	if m.Year < 1900 || m.Year > nextYear {
		return fmt.Errorf("invalid year")
	}
	return nil
}

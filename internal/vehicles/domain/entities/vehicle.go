package entities

import (
	"fmt"
	"github.com/google/uuid"
)

type Vehicle struct {
	ID    uuid.UUID
	Model Model
	Price float64
}

func (v *Vehicle) IsValid() error {
	if err := v.Model.IsValid(); err != nil {
		return fmt.Errorf("invalid car model. Reason: %v", err)
	}
	if v.Price < 0 {
		return fmt.Errorf("invalid price")
	}
	return nil
}

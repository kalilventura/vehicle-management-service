package valueobjects

import (
	"errors"

	"github.com/kalilventura/vehicle-management/internal/shared/domain"
)

type DoorsProps struct {
	Value int
}

// Doors represents the number of doors in a vehicle
type Doors struct {
	domain.BaseValueObject[DoorsProps]
}

// NewDoors creates a new Doors value object
func NewDoors(value int) (Doors, error) {
	if value < 2 || value > 5 {
		return Doors{}, errors.New("door count must be between 2 and 5")
	}
	return Doors{
		BaseValueObject: domain.NewBaseValueObject(DoorsProps{Value: value}),
	}, nil
}

// Value returns the doors value
func (d Doors) Value() int {
	return d.Props().Value
}


package valueobjects

import (
	"errors"

	"github.com/kalilventura/vehicle-management/internal/shared/domain"
)

type MileageProps struct {
	Value int
}

// Mileage represents vehicle mileage
type Mileage struct {
	domain.BaseValueObject[MileageProps]
}

// NewMileage creates a new Mileage value object
func NewMileage(value int) (Mileage, error) {
	if value < 0 {
		return Mileage{}, errors.New("mileage cannot be negative")
	}
	return Mileage{
		BaseValueObject: domain.NewBaseValueObject(MileageProps{Value: value}),
	}, nil
}

// Value returns the mileage value
func (m Mileage) Value() int {
	return m.Props().Value
}

// IsGreaterThan checks if mileage is greater than a value
func (m Mileage) IsGreaterThan(value int) bool {
	return m.Value() > value
}


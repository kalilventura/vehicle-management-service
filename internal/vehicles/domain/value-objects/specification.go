package valueobjects

import (
	"github.com/kalilventura/vehicle-management/internal/shared/domain"
)

type SpecificationProps struct {
	BodyType     BodyType
	Transmission Transmission
	FuelType     FuelType
	Mileage      Mileage
	Doors        Doors
	Engine       string
}

// Specification represents vehicle specifications
type Specification struct {
	domain.BaseValueObject[SpecificationProps]
}

// NewSpecification creates a new Specification value object
func NewSpecification(props SpecificationProps) Specification {
	return Specification{
		BaseValueObject: domain.NewBaseValueObject(props),
	}
}

// BodyType returns the body type
func (s Specification) BodyType() BodyType {
	return s.Props().BodyType
}

// Transmission returns the transmission
func (s Specification) Transmission() Transmission {
	return s.Props().Transmission
}

// FuelType returns the fuel type
func (s Specification) FuelType() FuelType {
	return s.Props().FuelType
}

// Mileage returns the mileage
func (s Specification) Mileage() Mileage {
	return s.Props().Mileage
}

// Doors returns the doors
func (s Specification) Doors() Doors {
	return s.Props().Doors
}

// Engine returns the engine description
func (s Specification) Engine() string {
	return s.Props().Engine
}

// MileageIsGreaterThan checks if mileage is greater than a value
func (s Specification) MileageIsGreaterThan(value int) bool {
	return s.Mileage().IsGreaterThan(value)
}


package entities

import "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"

type Specification struct {
	BodyType     dtos.BodyType
	Transmission dtos.Transmission
	FuelType     dtos.FuelType
	Mileage      dtos.Mileage
	Doors        dtos.Doors
	Engine       string
}

func (spec Specification) MileageIsGreaterThan(value int) bool {
	return spec.Mileage.Value() > value
}

func (spec Specification) GetBodyType() string {
	return spec.BodyType.Value()
}

func (spec Specification) GetFuelType() string {
	return spec.FuelType.Value()
}

func (spec Specification) GetMileage() int {
	return spec.Mileage.Value()
}

func (spec Specification) GetDoors() int {
	return spec.Doors.Value()
}

func (spec Specification) GetEngine() string {
	return spec.Engine
}

func (spec Specification) GetTransmission() string {
	return spec.Transmission.Value()
}

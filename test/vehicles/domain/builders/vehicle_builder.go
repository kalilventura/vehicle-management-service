package builders

import (
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
)

type VehicleBuilder struct{}

func NewVehicleBuilder() *VehicleBuilder {
	return &VehicleBuilder{}
}

func (b *VehicleBuilder) BuildValid() *entities.Vehicle {
	mileage, _ := dtos.NewMileage(0)
	specification := entities.Specification{
		Mileage: mileage,
	}
	return &entities.Vehicle{
		Condition:     dtos.New,
		Specification: specification,
	}
}

func (b *VehicleBuilder) BuildInvalid() *entities.Vehicle {
	mileage, _ := dtos.NewMileage(10)
	specification := entities.Specification{
		Mileage: mileage,
	}
	return &entities.Vehicle{
		Condition:     dtos.New,
		Specification: specification,
	}
}

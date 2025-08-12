package builders

import (
	global "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/kalilventura/vehicle-management/test/shared/domain/builders"
)

type VehicleBuilder struct {
	builders.BaseBuilder[entities.Vehicle]
}

func NewVehicleBuilder() *VehicleBuilder {
	return &VehicleBuilder{}
}

func (b *VehicleBuilder) WithSpecification(value entities.Specification) *VehicleBuilder {
	b.AppendModifier(func(e *entities.Vehicle) {
		e.Specification = value
	})
	return b
}

func (b *VehicleBuilder) WithCondition(value dtos.Condition) *VehicleBuilder {
	b.AppendModifier(func(e *entities.Vehicle) {
		e.Condition = value
	})
	return b
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

func (b *VehicleBuilder) BuildPagination() *global.PaginatedEntity[entities.Vehicle] {
	pagination := global.Pagination{}
	var list []entities.Vehicle

	for range 10 {
		list = append(list, *b.BuildValid())
	}

	page := global.NewPaginatedEntity(list, pagination)
	return &page
}

package builders

import (
  global "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
  "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
  valueobjects "github.com/kalilventura/vehicle-management/internal/vehicles/domain/value-objects"
  "github.com/kalilventura/vehicle-management/test/shared/domain/builders"
)

type VehicleBuilder struct {
  builders.BaseBuilder[entities.Vehicle]
}

func NewVehicleBuilder() *VehicleBuilder {
  return &VehicleBuilder{}
}

func (b *VehicleBuilder) WithYear(value int) *VehicleBuilder {
  year, _ := valueobjects.NewYear(value)
  b.AppendModifier(func(e *entities.Vehicle) {
    e.Year = year
  })
  return b
}

func (b *VehicleBuilder) WithSpecification(value entities.Specification) *VehicleBuilder {
  b.AppendModifier(func(e *entities.Vehicle) {
    e.Specification = value
  })
  return b
}

func (b *VehicleBuilder) WithCondition(value valueobjects.VehicleCondition) *VehicleBuilder {
  b.AppendModifier(func(e *entities.Vehicle) {
    e.Condition = value
  })
  return b
}

func (b *VehicleBuilder) BuildValid() *entities.Vehicle {
  mileage, _ := valueobjects.NewMileage(0)
  specification := entities.Specification{
    Mileage: mileage,
  }
  return &entities.Vehicle{
    Condition:     valueobjects.ConditionNew,
    Specification: specification,
  }
}

func (b *VehicleBuilder) BuildInvalid() *entities.Vehicle {
  mileage, _ := valueobjects.NewMileage(10)
  specification := entities.Specification{
    Mileage: mileage,
  }
  return &entities.Vehicle{
    Condition:     valueobjects.ConditionNew,
    Specification: specification,
  }
}

func (b *VehicleBuilder) BuildPagination() *global.PaginatedEntity[entities.Vehicle] {
  list := b.BuildMany()
  pagination := global.Pagination{}
  page := global.NewPaginatedEntity(list, pagination)
  return &page
}

package builders

import (
	"github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
)

type ListVehicleInputBuilder struct{}

func NewListVehicleInputBuilder() *ListVehicleInputBuilder {
	return &ListVehicleInputBuilder{}
}

func (b ListVehicleInputBuilder) Build() dtos.ListVehiclesInput {
	return dtos.ListVehiclesInput{
		Status:     nil,
		MinPrice:   nil,
		MaxPrice:   nil,
		Pagination: entities.Pagination{},
		SortBy:     "",
		SortOrder:  "",
	}
}

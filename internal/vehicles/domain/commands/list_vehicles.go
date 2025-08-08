package commands

import (
	global "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
)

type ListVehicles interface {
	Execute(input dtos.ListVehiclesInput, listeners ListVehiclesListeners)
}

type ListVehiclesListeners struct {
	OnSuccess             func(vehicles *global.PaginatedEntity[entities.Vehicle])
	OnInternalServerError func(err error)
}

package commands

import (
	global "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/commands"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
)

type ListVehiclesCommandStub struct {
	callback func(listeners commands.ListVehiclesListeners)
}

func NewListVehiclesCommandStub() *ListVehiclesCommandStub {
	return &ListVehiclesCommandStub{}
}

func (cmd *ListVehiclesCommandStub) WithOnSuccess(vehicles *global.PaginatedEntity[entities.Vehicle]) *ListVehiclesCommandStub {
	cmd.callback = func(listeners commands.ListVehiclesListeners) {
		listeners.OnSuccess(vehicles)
	}
	return cmd
}

func (cmd *ListVehiclesCommandStub) WithOnInternalServerError() *ListVehiclesCommandStub {
	cmd.callback = func(listeners commands.ListVehiclesListeners) {
		listeners.OnInternalServerError(getVehicleErr)
	}
	return cmd
}

func (cmd *ListVehiclesCommandStub) Execute(_ dtos.ListVehiclesInput, listeners commands.ListVehiclesListeners) {
	cmd.callback(listeners)
}

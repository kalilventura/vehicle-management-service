package commands

import (
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/commands"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
)

type UpdateVehicleCommandStub struct {
	callback func(listeners commands.UpdateVehicleListeners)
}

func NewUpdateVehicleCommandStub() *UpdateVehicleCommandStub {
	return &UpdateVehicleCommandStub{}
}

func (cmd *UpdateVehicleCommandStub) WithOnSuccess(vehicle entities.UpdateVehicleInput) *UpdateVehicleCommandStub {
	cmd.callback = func(listeners commands.UpdateVehicleListeners) {
		listeners.OnSuccess(&vehicle)
	}
	return cmd
}

func (cmd *UpdateVehicleCommandStub) WithOnNotFound() *UpdateVehicleCommandStub {
	cmd.callback = func(listeners commands.UpdateVehicleListeners) {
		listeners.OnNotFound()
	}
	return cmd
}

func (cmd *UpdateVehicleCommandStub) WithOnInternalServerError() *UpdateVehicleCommandStub {
	cmd.callback = func(listeners commands.UpdateVehicleListeners) {
		listeners.OnInternalServerError(getVehicleErr)
	}
	return cmd
}

func (cmd *UpdateVehicleCommandStub) Execute(_ *entities.UpdateVehicleInput, listeners commands.UpdateVehicleListeners) {
	cmd.callback(listeners)
}

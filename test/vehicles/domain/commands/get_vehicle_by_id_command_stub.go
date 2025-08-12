package commands

import (
	"errors"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/commands"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
)

var getVehicleErr = errors.New("getVehicleErr")

type GetVehicleByIDCommandStub struct {
	callback func(listeners commands.GetVehicleByIDListeners)
}

func NewGetVehicleByIDCommandStub() *GetVehicleByIDCommandStub {
	return &GetVehicleByIDCommandStub{}
}

func (cmd *GetVehicleByIDCommandStub) WithOnSuccess(vehicle entities.Vehicle) *GetVehicleByIDCommandStub {
	cmd.callback = func(listeners commands.GetVehicleByIDListeners) {
		listeners.OnSuccess(&vehicle)
	}
	return cmd
}

func (cmd *GetVehicleByIDCommandStub) WithOnNotFound() *GetVehicleByIDCommandStub {
	cmd.callback = func(listeners commands.GetVehicleByIDListeners) {
		listeners.OnNotFound()
	}
	return cmd
}

func (cmd *GetVehicleByIDCommandStub) WithOnInternalServerError() *GetVehicleByIDCommandStub {
	cmd.callback = func(listeners commands.GetVehicleByIDListeners) {
		listeners.OnInternalServerError(getVehicleErr)
	}
	return cmd
}

func (cmd *GetVehicleByIDCommandStub) Execute(_ string, listeners commands.GetVehicleByIDListeners) {
	cmd.callback(listeners)
}

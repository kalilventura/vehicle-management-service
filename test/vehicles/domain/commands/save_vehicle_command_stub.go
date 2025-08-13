package commands

import (
  "errors"

  "github.com/kalilventura/vehicle-management/internal/vehicles/domain/commands"
  "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
)

var saveVehicleErr = errors.New("getVehicleErr")

type SaveVehicleCommandStub struct {
  callback func(listeners commands.SaveVehicleListeners)
}

func NewSaveVehicleCommandStub() *SaveVehicleCommandStub {
  return &SaveVehicleCommandStub{}
}

func (cmd *SaveVehicleCommandStub) WithOnSuccess(vehicle entities.Vehicle) *SaveVehicleCommandStub {
  cmd.callback = func(listeners commands.SaveVehicleListeners) {
    listeners.OnSuccess(&vehicle)
  }
  return cmd
}

func (cmd *SaveVehicleCommandStub) WithOnNotValid() *SaveVehicleCommandStub {
  cmd.callback = func(listeners commands.SaveVehicleListeners) {
    listeners.OnNotValid(saveVehicleErr)
  }
  return cmd
}

func (cmd *SaveVehicleCommandStub) WithOnInternalServerError() *SaveVehicleCommandStub {
  cmd.callback = func(listeners commands.SaveVehicleListeners) {
    listeners.OnInternalServerError(getVehicleErr)
  }
  return cmd
}

func (cmd *SaveVehicleCommandStub) Execute(_ *entities.Vehicle, listeners commands.SaveVehicleListeners) {
  cmd.callback(listeners)
}

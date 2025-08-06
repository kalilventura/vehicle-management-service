package commands

import (
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/repositories"
)

type SaveVehicleCommand struct {
	repository repositories.VehiclesRepository
}

func NewSaveVehicleCommand(repository repositories.VehiclesRepository) *SaveVehicleCommand {
	return &SaveVehicleCommand{repository}
}

func (cmd *SaveVehicleCommand) Execute(vehicle *entities.Vehicle, listeners SaveVehicleListeners) {
	if err := vehicle.IsValid(); err != nil {
		listeners.OnNotValid(err)
		return
	}

	if err := cmd.repository.Save(vehicle); err != nil {
		listeners.OnInternalServerError(err)
		return
	}
	listeners.OnSuccess(vehicle)
}

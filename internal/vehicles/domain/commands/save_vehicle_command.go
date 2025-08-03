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

func (cmd *SaveVehicleCommand) Execute(vehicle *entities.Vehicle) {
	if err := vehicle.IsValid(); err != nil {
		return
	}
	if err := cmd.repository.Save(vehicle); err != nil {
		return
	}
}

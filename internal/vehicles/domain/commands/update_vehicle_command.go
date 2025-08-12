package commands

import (
	"errors"

	domainerror "github.com/kalilventura/vehicle-management/internal/shared/domain/errors"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/repositories"
)

type UpdateVehicleCommand struct {
	repository repositories.VehiclesRepository
}

func NewUpdateVehicleCommand(
	repository repositories.VehiclesRepository) *UpdateVehicleCommand {
	return &UpdateVehicleCommand{repository}
}

func (cmd *UpdateVehicleCommand) Execute(input *entities.UpdateVehicleInput, listeners UpdateVehicleListeners) {
	_, err := cmd.repository.GetByID(input.ID)
	switch {
	case errors.Is(err, domainerror.ErrRecordNotFound):
		listeners.OnNotFound()
		return
	case err != nil:
		listeners.OnInternalServerError(err)
		return
	}
	if updateErr := cmd.repository.Update(input); updateErr != nil {
		listeners.OnInternalServerError(updateErr)
		return
	}
	listeners.OnSuccess(input)
}

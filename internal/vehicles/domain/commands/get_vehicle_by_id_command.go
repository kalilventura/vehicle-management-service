package commands

import (
	"errors"

	domainerror "github.com/kalilventura/vehicle-management/internal/shared/domain/errors"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/repositories"
)

type GetVehicleByIDCommand struct {
	repository repositories.VehiclesRepository
}

func NewGetVehicleByIDCommand(repository repositories.VehiclesRepository) *GetVehicleByIDCommand {
	return &GetVehicleByIDCommand{repository}
}

func (cmd *GetVehicleByIDCommand) Execute(ID string, listeners GetVehicleByIDListeners) {
	target, err := cmd.repository.GetByID(ID)
	switch {
	case errors.Is(err, domainerror.ErrRecordNotFound):
		listeners.OnNotFound()
		return
	case err != nil:
		listeners.OnInternalServerError(err)
		return
	}
	listeners.OnSuccess(target)
}

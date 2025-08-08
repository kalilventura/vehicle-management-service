package commands

import (
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/repositories"
)

type ListVehiclesCommand struct {
	repository repositories.VehiclesRepository
}

func NewListVehiclesCommand(repository repositories.VehiclesRepository) *ListVehiclesCommand {
	return &ListVehiclesCommand{repository}
}

func (cmd *ListVehiclesCommand) Execute(input dtos.ListVehiclesInput, listeners ListVehiclesListeners) {
	response, err := cmd.repository.FindWithFilters(input)
	if err != nil {
		listeners.OnInternalServerError(err)
		return
	}
	listeners.OnSuccess(response)
}

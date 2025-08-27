package commands

import (
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/repositories"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/services"
)

type SellVehicleCommand struct {
	payment    services.PaymentsService
	repository repositories.VehiclesRepository
}

func NewSellVehicleCommand(
	payment services.PaymentsService,
	repository repositories.VehiclesRepository) *SellVehicleCommand {
	return &SellVehicleCommand{payment, repository}
}

func (cmd *SellVehicleCommand) Execute(inputRequest *entities.SellVehicle, listeners SellVehicleListeners) {
	serviceListeners := services.PaymentsServiceListeners{
		OnSuccess: func(sellRequest *entities.SellVehicle) {
			cmd.updateVehicle(sellRequest, listeners)
		},
		OnBadRequest:          listeners.OnBadRequest,
		OnInternalServerError: listeners.OnInternalServerError,
	}
	cmd.payment.Pay(inputRequest, serviceListeners)
}

func (cmd *SellVehicleCommand) updateVehicle(sell *entities.SellVehicle, listeners SellVehicleListeners) {
	newStatus := dtos.Sold
	update := &entities.UpdateVehicleInput{
		ID:     sell.VehicleID,
		Status: &newStatus,
	}
	if err := cmd.repository.Update(update); err != nil {
		listeners.OnInternalServerError(err)
		return
	}
	listeners.OnSuccess(sell)
}

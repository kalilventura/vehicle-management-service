package requests

import "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"

// SellVehicleRequest
// @Description Object that represents a sell
type SellVehicleRequest struct {
	CPF    string  `json:"cpf" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
} // @name SellVehicleRequest

func (dto SellVehicleRequest) ToDomain(vehicleID string) *entities.SellVehicle {
	return &entities.SellVehicle{
		VehicleID: vehicleID,
		Cpf:       dto.CPF,
		Amount:    dto.Amount,
	}
}

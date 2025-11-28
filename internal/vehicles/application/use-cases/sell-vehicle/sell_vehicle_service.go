package sellvehicle

import (
	"errors"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/repositories"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/services"
)

// SellVehicleService is the application service for selling a vehicle
type SellVehicleService struct {
	repository     repositories.VehiclesRepository
	paymentsService services.PaymentsService
}

// NewSellVehicleService creates a new SellVehicleService
func NewSellVehicleService(
	repository repositories.VehiclesRepository,
	paymentsService services.PaymentsService,
) *SellVehicleService {
	return &SellVehicleService{
		repository:     repository,
		paymentsService: paymentsService,
	}
}

// Execute executes the sell vehicle use case
func (s *SellVehicleService) Execute(dto SellVehicleDTO) error {
	// Get vehicle
	vehicle, err := s.repository.GetByID(dto.VehicleID)
	if err != nil {
		return err
	}

	if vehicle == nil {
		return errors.New("vehicle not found")
	}

	// Process payment
	if err := s.paymentsService.ProcessPayment(dto.CPF, dto.Amount); err != nil {
		return err
	}

	// Sell vehicle (domain logic)
	if err := vehicle.Sell(); err != nil {
		return err
	}

	// Save vehicle
	if err := s.repository.Save(vehicle); err != nil {
		return err
	}

	// Note: Domain events should be published here by the infrastructure layer
	// after the transaction is committed

	return nil
}


package sellvehicle

import (
	"errors"
	"strings"

	domainerr "github.com/kalilventura/vehicle-management/internal/shared/domain/errors"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/exceptions"
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
		if errors.Is(err, domainerr.ErrRecordNotFound) {
			return exceptions.NewVehicleNotFoundException(dto.VehicleID)
		}
		return err
	}

	if vehicle == nil {
		return exceptions.NewVehicleNotFoundException(dto.VehicleID)
	}

	// Process payment
	if err := s.paymentsService.ProcessPayment(dto.CPF, dto.Amount); err != nil {
		// Check if it's a bad request error
		if strings.Contains(err.Error(), "bad request") {
			return exceptions.NewPaymentException(dto.CPF, dto.Amount, "invalid payment data")
		}
		// For other payment errors, return as-is (will be handled as 500)
		return err
	}

	// Sell vehicle (domain logic)
	if err := vehicle.Sell(); err != nil {
		// Check if it's a status error
		if strings.Contains(err.Error(), "cannot be sold") {
			return exceptions.NewInvalidVehicleStatusException(
				dto.VehicleID,
				vehicle.Status().Value(),
				"sell",
			)
		}
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


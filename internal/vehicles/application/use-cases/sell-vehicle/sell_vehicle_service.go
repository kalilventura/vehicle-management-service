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
	repository      repositories.VehiclesRepository
	paymentsService services.PaymentsService
}

// NewSellVehicleService creates a new SellVehicleService
func NewSellVehicleService(
	repository repositories.VehiclesRepository,
	paymentsService services.PaymentsService,
) *SellVehicleService {
	return &SellVehicleService{
		repository:      repository,
		paymentsService: paymentsService,
	}
}

// Execute executes the sell vehicle use case
func (s *SellVehicleService) Execute(dto SellVehicleDTO) error {
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
	if paymentErr := s.paymentsService.ProcessPayment(dto.CPF, dto.Amount); paymentErr != nil {
		if strings.Contains(paymentErr.Error(), "bad request") {
			return exceptions.NewPaymentException(dto.CPF, dto.Amount, "invalid payment data")
		}
		return paymentErr
	}

	if sellErr := vehicle.Sell(); sellErr != nil {
		if strings.Contains(sellErr.Error(), "cannot be sold") {
			return exceptions.NewInvalidVehicleStatusException(
				dto.VehicleID,
				vehicle.Status().Value(),
				"sell",
			)
		}
		return sellErr
	}

	if saveErr := s.repository.Save(vehicle); saveErr != nil {
		return saveErr
	}

	// Note: Domain events should be published here by the infrastructure layer
	// after the transaction is committed

	return nil
}

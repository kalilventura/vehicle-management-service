package getvehicle

import (
	"errors"

	domainerr "github.com/kalilventura/vehicle-management/internal/shared/domain/errors"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/dtos"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/factories"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/exceptions"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/repositories"
)

// GetVehicleService is the application service for getting a vehicle by ID
type GetVehicleService struct {
	repository repositories.VehiclesRepository
}

// NewGetVehicleService creates a new GetVehicleService
func NewGetVehicleService(
	repository repositories.VehiclesRepository,
) *GetVehicleService {
	return &GetVehicleService{
		repository: repository,
	}
}

// Execute executes the get vehicle use case
func (s *GetVehicleService) Execute(vehicleID string) (dtos.VehicleResponseDTO, error) {
	vehicle, err := s.repository.GetByID(vehicleID)
	if err != nil {
		if errors.Is(err, domainerr.ErrRecordNotFound) {
			return dtos.VehicleResponseDTO{}, exceptions.NewVehicleNotFoundException(vehicleID)
		}
		return dtos.VehicleResponseDTO{}, err
	}

	if vehicle == nil {
		return dtos.VehicleResponseDTO{}, exceptions.NewVehicleNotFoundException(vehicleID)
	}
	return factories.ToResponseDTO(vehicle), nil
}

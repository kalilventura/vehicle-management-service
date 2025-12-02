package updatevehicle

import (
	"errors"

	domainerr "github.com/kalilventura/vehicle-management/internal/shared/domain/errors"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/dtos"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/factories"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/exceptions"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/repositories"
)

// UpdateVehicleService is the application service for updating vehicles
type UpdateVehicleService struct {
	repository repositories.VehiclesRepository
}

// NewUpdateVehicleService creates a new UpdateVehicleService
func NewUpdateVehicleService(
	repository repositories.VehiclesRepository,
) *UpdateVehicleService {
	return &UpdateVehicleService{
		repository: repository,
	}
}

// Execute executes the update vehicle use case
func (s *UpdateVehicleService) Execute(input *entities.UpdateVehicleInput) (dtos.VehicleResponseDTO, error) {
	// Check if a vehicle exists
	_, err := s.repository.GetByID(input.ID)
	if err != nil {
		if errors.Is(err, domainerr.ErrRecordNotFound) {
			return dtos.VehicleResponseDTO{}, exceptions.NewVehicleNotFoundException(input.ID)
		}
		return dtos.VehicleResponseDTO{}, err
	}

	// Update vehicle
	if err := s.repository.Update(input); err != nil {
		return dtos.VehicleResponseDTO{}, err
	}

	// Get an updated vehicle
	vehicle, err := s.repository.GetByID(input.ID)
	if err != nil {
		return dtos.VehicleResponseDTO{}, err
	}

	// Convert to response DTO
	return factories.ToResponseDTO(vehicle), nil
}

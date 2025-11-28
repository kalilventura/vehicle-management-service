package getvehicle

import (
	domainerr "github.com/kalilventura/vehicle-management/internal/shared/domain/errors"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/dtos"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/mappers"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/exceptions"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/repositories"
)

// GetVehicleService is the application service for getting a vehicle by ID
type GetVehicleService struct {
	repository repositories.VehiclesRepository
	mapper      *mappers.VehicleMapper
}

// NewGetVehicleService creates a new GetVehicleService
func NewGetVehicleService(
	repository repositories.VehiclesRepository,
	mapper *mappers.VehicleMapper,
) *GetVehicleService {
	return &GetVehicleService{
		repository: repository,
		mapper:      mapper,
	}
}

// Execute executes the get vehicle use case
func (s *GetVehicleService) Execute(vehicleID string) (dtos.VehicleResponseDTO, error) {
	vehicle, err := s.repository.GetByID(vehicleID)
	if err != nil {
		// Map repository errors to domain exceptions
		if err == domainerr.ErrRecordNotFound {
			return dtos.VehicleResponseDTO{}, exceptions.NewVehicleNotFoundException(vehicleID)
		}
		return dtos.VehicleResponseDTO{}, err
	}

	if vehicle == nil {
		return dtos.VehicleResponseDTO{}, exceptions.NewVehicleNotFoundException(vehicleID)
	}

	return s.mapper.ToResponseDTO(vehicle), nil
}


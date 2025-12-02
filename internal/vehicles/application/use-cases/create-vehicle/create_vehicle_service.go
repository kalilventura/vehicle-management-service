package createvehicle

import (
  "github.com/kalilventura/vehicle-management/internal/vehicles/application/dtos"
  "github.com/kalilventura/vehicle-management/internal/vehicles/application/factories"
  "github.com/kalilventura/vehicle-management/internal/vehicles/domain/repositories"
)

// CreateVehicleService is the application service for creating vehicles
type CreateVehicleService struct {
  repository repositories.VehiclesRepository
}

// NewCreateVehicleService creates a new CreateVehicleService
func NewCreateVehicleService(
  repository repositories.VehiclesRepository,
) *CreateVehicleService {
  return &CreateVehicleService{
    repository: repository,
  }
}

// Execute executes the creation vehicle use case
func (s *CreateVehicleService) Execute(dto dtos.CreateVehicleDTO) (dtos.VehicleResponseDTO, error) {
  // Convert DTO to domain entity
  vehicle, err := factories.ToDomainEntity(dto)
  if err != nil {
    return dtos.VehicleResponseDTO{}, err
  }

  // Save vehicle
  if err := s.repository.Save(vehicle); err != nil {
    return dtos.VehicleResponseDTO{}, err
  }

  // Convert to response DTO
  response := factories.ToResponseDTO(vehicle)

  // Note: Domain events should be published here by the infrastructure layer
  // after the transaction is committed
  // The domain events can be accessed via vehicle.DomainEvents()

  return response, nil
}

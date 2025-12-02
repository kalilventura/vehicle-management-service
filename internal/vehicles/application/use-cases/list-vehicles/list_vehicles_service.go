package listvehicles

import (
	global "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	appdtos "github.com/kalilventura/vehicle-management/internal/vehicles/application/dtos"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/factories"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/repositories"
	valueobjects "github.com/kalilventura/vehicle-management/internal/vehicles/domain/value-objects"
)

// ListVehiclesService is the application service for listing vehicles
type ListVehiclesService struct {
	repository repositories.VehiclesRepository
}

// NewListVehiclesService creates a new ListVehiclesService
func NewListVehiclesService(
	repository repositories.VehiclesRepository,
) *ListVehiclesService {
	return &ListVehiclesService{
		repository: repository,
	}
}

// Execute executes the list of vehicles use case
func (s *ListVehiclesService) Execute(
	input valueobjects.ListVehiclesCriteria) (*global.PaginatedEntity[appdtos.VehicleResponseDTO], error) {
	vehicles, err := s.repository.FindWithFilters(input)
	if err != nil {
		return nil, err
	}

	var responseList []appdtos.VehicleResponseDTO
	for _, vehicle := range vehicles.Content {
		responseList = append(responseList, factories.ToResponseDTO(&vehicle))
	}

	page := global.NewPaginatedEntity(responseList, vehicles.Pagination)
	return &page, nil
}

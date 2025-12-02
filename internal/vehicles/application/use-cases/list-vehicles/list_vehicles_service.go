package listvehicles

import (
	global "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	appdtos "github.com/kalilventura/vehicle-management/internal/vehicles/application/dtos"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/mappers"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/repositories"
)

// ListVehiclesService is the application service for listing vehicles
type ListVehiclesService struct {
	repository repositories.VehiclesRepository
	mapper     *mappers.VehicleMapper
}

// NewListVehiclesService creates a new ListVehiclesService
func NewListVehiclesService(
	repository repositories.VehiclesRepository,
	mapper *mappers.VehicleMapper,
) *ListVehiclesService {
	return &ListVehiclesService{
		repository: repository,
		mapper:     mapper,
	}
}

// Execute executes the list vehicles use case
func (s *ListVehiclesService) Execute(input dtos.ListVehiclesInput) (*global.PaginatedEntity[appdtos.VehicleResponseDTO], error) {
	vehicles, err := s.repository.FindWithFilters(input)
	if err != nil {
		return nil, err
	}

	var responseList []appdtos.VehicleResponseDTO
	for _, vehicle := range vehicles.Content {
		responseList = append(responseList, s.mapper.ToResponseDTO(&vehicle))
	}

	return &global.PaginatedEntity[appdtos.VehicleResponseDTO]{
		Content:    responseList,
		Pagination: vehicles.Pagination,
	}, nil
}

package repositories

import (
	global "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
)

type VehiclesRepository interface {
	Save(vehicle *entities.Vehicle) error
	GetByID(ID string) (*entities.Vehicle, error)
	FindWithFilters(input dtos.ListVehiclesInput) (*global.PaginatedEntity[entities.Vehicle], error)
}

package repositories

import (
	global "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	valueobjects "github.com/kalilventura/vehicle-management/internal/vehicles/domain/value-objects"
)

type VehiclesRepository interface {
	Save(vehicle *entities.Vehicle) error
	GetByID(ID string) (*entities.Vehicle, error)
	FindWithFilters(input valueobjects.ListVehiclesCriteria) (*global.PaginatedEntity[entities.Vehicle], error)
}

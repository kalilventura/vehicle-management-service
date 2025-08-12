package repositories

import (
	global "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
)

type InMemoryVehiclesRepository struct {
	err      error
	vehicles *global.PaginatedEntity[entities.Vehicle]
}

func NewInMemoryVehiclesRepository() *InMemoryVehiclesRepository {
	return &InMemoryVehiclesRepository{}
}

func (r *InMemoryVehiclesRepository) WithError(err error) *InMemoryVehiclesRepository {
	r.err = err
	return r
}

func (r *InMemoryVehiclesRepository) WithVehicles(
	vehicles *global.PaginatedEntity[entities.Vehicle]) *InMemoryVehiclesRepository {
	r.vehicles = vehicles
	return r
}

func (r *InMemoryVehiclesRepository) Save(_ *entities.Vehicle) error {
	return r.err
}

func (r *InMemoryVehiclesRepository) Update(_ *entities.UpdateVehicleInput) error {
	return r.err
}

func (r *InMemoryVehiclesRepository) GetByID(_ string) (*entities.Vehicle, error) {
	if r.vehicles != nil {
		return &r.vehicles.Content[0], nil
	}
	return nil, r.err
}

func (r *InMemoryVehiclesRepository) FindWithFilters(
	_ dtos.ListVehiclesInput) (*global.PaginatedEntity[entities.Vehicle], error) {
	return r.vehicles, r.err
}

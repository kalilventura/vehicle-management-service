package repositories

import "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"

type InMemoryVehiclesRepository struct {
	err error
}

func NewInMemoryVehiclesRepository() *InMemoryVehiclesRepository {
	return &InMemoryVehiclesRepository{}
}

func (r *InMemoryVehiclesRepository) WithError(err error) *InMemoryVehiclesRepository {
	r.err = err
	return r
}

func (r *InMemoryVehiclesRepository) Save(_ *entities.Vehicle) error {
	return r.err
}

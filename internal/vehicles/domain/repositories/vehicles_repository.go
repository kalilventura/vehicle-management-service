package repositories

import "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"

type VehiclesRepository interface {
	GetByID(ID string) (*entities.Vehicle, error)
	Save(vehicle *entities.Vehicle) error
}

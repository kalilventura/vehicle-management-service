package repositories

import "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"

type VehiclesRepository interface {
	Save(vehicle *entities.Vehicle) error
}

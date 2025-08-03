package repositories

import (
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"gorm.io/gorm"
)

type GormVehiclesRepository struct {
	client *gorm.DB
}

func NewGormVehiclesRepository(client *gorm.DB) *GormVehiclesRepository {
	return &GormVehiclesRepository{client: client}
}

func (r *GormVehiclesRepository) Save(vehicle *entities.Vehicle) error {
	return nil
}

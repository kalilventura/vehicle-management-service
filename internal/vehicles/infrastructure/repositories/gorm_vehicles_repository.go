package repositories

import (
	"fmt"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/repositories/models"
	"gorm.io/gorm"
)

type GormVehiclesRepository struct {
	client *gorm.DB
}

func NewGormVehiclesRepository(client *gorm.DB) *GormVehiclesRepository {
	return &GormVehiclesRepository{client: client}
}

func (r *GormVehiclesRepository) Save(vehicle *entities.Vehicle) error {
	gormEntity := models.FromDomain(vehicle)
	if err := r.client.Save(&gormEntity).Error; err != nil {
		return fmt.Errorf("failed to save vehicle. Reason: %w", err)
	}
	vehicle.ID = gormEntity.ID
	return nil
}

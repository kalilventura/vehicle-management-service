package repositories

import (
  "errors"
  "fmt"

  domainerr "github.com/kalilventura/vehicle-management/internal/shared/domain/errors"
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

func (r *GormVehiclesRepository) GetByID(ID string) (*entities.Vehicle, error) {
  vehicle := &models.GormVehicle{}
  err := r.client.First(vehicle, "id = ?", ID).Error
  if err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
      return nil, domainerr.ErrRecordNotFound
    }
    return nil, fmt.Errorf("failed to find vehicle. Reason: %w", err)
  }
  return vehicle.ToDomain(), nil
}

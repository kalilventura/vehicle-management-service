package infrastructure

import (
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/repositories/models"
	"gorm.io/gorm"
)

func CreateDatabaseStructure(db *gorm.DB) error {
	err := db.AutoMigrate(models.GormVehicle{})
	return err
}

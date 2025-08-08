package mappers

import (
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/repositories/models"
)

func MapToDomainList(gormEntities []models.GormVehicle) []entities.Vehicle {
	list := make([]entities.Vehicle, len(gormEntities))
	for i, gormEntity := range gormEntities {
		mapped := gormEntity.ToDomain()
		list[i] = *mapped
	}
	return list
}

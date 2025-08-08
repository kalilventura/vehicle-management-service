package repositories

import (
	"errors"
	"fmt"

	global "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	domainerr "github.com/kalilventura/vehicle-management/internal/shared/domain/errors"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/repositories/mappers"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/repositories/models"
	"gorm.io/gorm"
)

type GormVehiclesRepository struct {
	client *gorm.DB
}

func NewGormVehiclesRepository(client *gorm.DB) *GormVehiclesRepository {
	return &GormVehiclesRepository{client: client}
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

func (r *GormVehiclesRepository) Save(vehicle *entities.Vehicle) error {
	gormEntity := models.FromDomain(vehicle)
	if err := r.client.Save(&gormEntity).Error; err != nil {
		return fmt.Errorf("failed to save vehicle. Reason: %w", err)
	}
	vehicle.ID = gormEntity.ID
	return nil
}

func (r *GormVehiclesRepository) FindWithFilters(
	filter dtos.ListVehiclesInput) (*global.PaginatedEntity[entities.Vehicle], error) {
	var list []models.GormVehicle
	query := r.client.Model(&models.GormVehicle{})

	if filter.Status != nil {
		query = query.Where("status = ?", filter.Status.Value())
	}
	if filter.MinPrice != nil {
		query = query.Where("min_price = ?", filter.MinPrice.Value())
	}
	if filter.MaxPrice != nil {
		query = query.Where("max_price = ?", filter.MaxPrice.Value())
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("failed to count vehicles. Reason: %w", err)
	}
	filter.Pagination.TotalElements = total

	if filter.SortBy != "" && filter.SortOrder != "" {
		query = query.Order(fmt.Sprintf("%s %s", filter.SortBy, filter.SortOrder))
	}

	err := query.
		Offset(filter.Pagination.Offset()).
		Limit(filter.Pagination.Size).
		Find(&list).
		Error
	if err != nil {
		return nil, fmt.Errorf("failed to find vehicles. Reason: %w", err)
	}

	entityList := mappers.MapToDomainList(list)
	pageResponse := global.NewPaginatedEntity(entityList, filter.Pagination)
	return &pageResponse, nil
}

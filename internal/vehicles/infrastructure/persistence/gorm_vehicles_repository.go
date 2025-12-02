package persistence

import (
	"errors"
	"fmt"

	global "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	domainerr "github.com/kalilventura/vehicle-management/internal/shared/domain/errors"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	valueobjects "github.com/kalilventura/vehicle-management/internal/vehicles/domain/value-objects"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/persistence/factories"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/persistence/models"
	"gorm.io/gorm"
)

// GormVehiclesRepository implements the VehiclesRepository interface using GORM
type GormVehiclesRepository struct {
	client *gorm.DB
}

// NewGormVehiclesRepository creates a new GormVehiclesRepository
func NewGormVehiclesRepository(client *gorm.DB) *GormVehiclesRepository {
	return &GormVehiclesRepository{
		client: client,
	}
}

// Save saves a vehicle
func (r *GormVehiclesRepository) Save(vehicle *entities.Vehicle) error {
	gormEntity := factories.ToOrmEntity(vehicle)
	if err := r.client.Save(gormEntity).Error; err != nil {
		return fmt.Errorf("failed to save vehicle: %w", err)
	}
	return nil
}

// GetByID gets a vehicle by ID
func (r *GormVehiclesRepository) GetByID(ID string) (*entities.Vehicle, error) {
	vehicle := &models.GormVehicle{}
	err := r.client.First(vehicle, "id = ?", ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerr.ErrRecordNotFound
		}
		return nil, fmt.Errorf("failed to find vehicle: %w", err)
	}
	return factories.ToDomainEntity(vehicle)
}

// FindWithFilters finds vehicles with filters
func (r *GormVehiclesRepository) FindWithFilters(
	filter valueobjects.ListVehiclesCriteria) (*global.PaginatedEntity[entities.Vehicle], error) {
	var list []models.GormVehicle
	query := r.client.Model(&models.GormVehicle{})

	if filter.Status != nil {
		query = query.Where("status = ?", filter.Status.Value())
	}
	if filter.MinPrice != nil {
		query = query.Where("price >= ?", filter.MinPrice.Amount())
	}
	if filter.MaxPrice != nil {
		query = query.Where("price <= ?", filter.MaxPrice.Amount())
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("failed to count vehicles: %w", err)
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
		return nil, fmt.Errorf("failed to find vehicles: %w", err)
	}

	// Convert to domain entities
	entityList := make([]entities.Vehicle, 0, len(list))
	for _, gormVehicle := range list {
		vehicle, facErr := factories.ToDomainEntity(&gormVehicle)
		if facErr != nil {
			return nil, fmt.Errorf("failed to convert vehicle: %w", facErr)
		}
		entityList = append(entityList, *vehicle)
	}

	pageResponse := global.NewPaginatedEntity(entityList, filter.Pagination)
	return &pageResponse, nil
}

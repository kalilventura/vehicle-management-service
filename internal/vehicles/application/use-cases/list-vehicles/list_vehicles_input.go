package listvehicles

import (
	"github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	valueobjects "github.com/kalilventura/vehicle-management/internal/vehicles/domain/value-objects"
)

type Input struct {
	Status     *valueobjects.VehicleStatus
	MinPrice   *valueobjects.Price
	MaxPrice   *valueobjects.Price
	Pagination entities.Pagination
	SortBy     string
	SortOrder  string
}

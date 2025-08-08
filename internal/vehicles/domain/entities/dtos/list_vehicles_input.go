package dtos

import (
	"github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
)

type ListVehiclesInput struct {
	Status     *Status
	MinPrice   *Price
	MaxPrice   *Price
	Pagination entities.Pagination
	SortBy     string
	SortOrder  string
}

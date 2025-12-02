package valueobjects

import (
  "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
)

// ListVehiclesCriteria represents the search criteria for listing vehicles
type ListVehiclesCriteria struct {
  Status     *VehicleStatus
  MinPrice   *Price
  MaxPrice   *Price
  Pagination entities.Pagination
  SortBy     string
  SortOrder  string
}

package entities

import "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"

type UpdateVehicleInput struct {
	ID          string
	Color       *string
	Description *string
	Price       *dtos.Price
	Features    *UpdateFeaturesInput
	Status      *dtos.Status
	Condition   *dtos.Condition
}

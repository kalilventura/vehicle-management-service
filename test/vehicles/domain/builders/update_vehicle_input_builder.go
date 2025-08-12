package builders

import "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"

type UpdateVehicleInputBuilder struct{}

func NewUpdateVehicleInputBuilder() *UpdateVehicleInputBuilder {
	return &UpdateVehicleInputBuilder{}
}

func (b *UpdateVehicleInputBuilder) BuildValid() *entities.UpdateVehicleInput {
	return &entities.UpdateVehicleInput{}
}

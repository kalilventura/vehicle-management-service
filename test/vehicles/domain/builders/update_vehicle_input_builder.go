package builders

import (
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/test/shared/domain/builders"
)

type UpdateVehicleInputBuilder struct {
	builders.BaseBuilder[entities.UpdateVehicleInput]
}

func NewUpdateVehicleInputBuilder() *UpdateVehicleInputBuilder {
	return &UpdateVehicleInputBuilder{}
}

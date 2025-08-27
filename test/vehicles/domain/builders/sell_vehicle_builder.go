package builders

import (
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/test/shared/domain/builders"
)

type SellVehicleBuilder struct {
	builders.BaseBuilder[entities.SellVehicle]
}

func NewSellVehicleBuilder() *SellVehicleBuilder {
	return &SellVehicleBuilder{}
}

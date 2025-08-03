package repositories

import (
	"github.com/google/wire"
	iface "github.com/kalilventura/vehicle-management/internal/vehicles/domain/repositories"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	NewGormVehiclesRepository,
	wire.Bind(new(iface.VehiclesRepository), new(*GormVehiclesRepository)),
)

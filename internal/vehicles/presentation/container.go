package presentation

import (
	"github.com/google/wire"
	"github.com/kalilventura/vehicle-management/internal/vehicles/presentation/filters"
	"github.com/kalilventura/vehicle-management/internal/vehicles/presentation/mappers"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	filters.NewVehicleExceptionFilter,
	mappers.NewVehicleResponseMapper,
)


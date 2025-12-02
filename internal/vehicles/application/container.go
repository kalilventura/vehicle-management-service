package application

import (
	"github.com/google/wire"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/mappers"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/use-cases/create-vehicle"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/use-cases/get-vehicle"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/use-cases/list-vehicles"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/use-cases/sell-vehicle"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/use-cases/update-vehicle"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	mappers.NewVehicleMapper,
	createvehicle.NewCreateVehicleService,
	getvehicle.NewGetVehicleService,
	listvehicles.NewListVehiclesService,
	updatevehicle.NewUpdateVehicleService,
	sellvehicle.NewSellVehicleService,
)


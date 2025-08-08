package commands

import "github.com/google/wire"

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	NewSaveVehicleCommand,
	wire.Bind(new(SaveVehicle), new(*SaveVehicleCommand)),
	NewGetVehicleByIDCommand,
	wire.Bind(new(GetVehicleByID), new(*GetVehicleByIDCommand)),
	NewListVehiclesCommand,
	wire.Bind(new(ListVehicles), new(*ListVehiclesCommand)),
)

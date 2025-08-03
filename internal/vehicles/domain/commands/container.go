package commands

import (
	"github.com/google/wire"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	NewSaveVehicleCommand,
	wire.Bind(new(SaveVehicle), new(*SaveVehicleCommand)),
)

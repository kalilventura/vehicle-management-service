package vehicles

import (
	"github.com/google/wire"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/commands"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/repositories"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	commands.Container,
	repositories.Container,
	controllers.Container,
	NewModule,
)

package vehicles

import (
	"github.com/google/wire"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/commands"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/repositories"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/services"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	commands.Container,
	repositories.Container,
	services.Container,
	controllers.Container,
	NewModule,
)

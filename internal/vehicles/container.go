package vehicles

import (
	"github.com/google/wire"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	controllers.Container,
	NewModule,
)

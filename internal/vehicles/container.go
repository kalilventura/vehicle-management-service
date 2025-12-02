package vehicles

import (
	"github.com/google/wire"
	appcontainer "github.com/kalilventura/vehicle-management/internal/vehicles/application"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/repositories"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/services"
	presentationcontainer "github.com/kalilventura/vehicle-management/internal/vehicles/presentation"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	persistence.Container,
	services.Container,
	appcontainer.Container,
	presentationcontainer.Container,
	NewModule,
)

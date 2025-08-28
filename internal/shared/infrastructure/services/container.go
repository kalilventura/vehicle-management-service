package services

import (
	"github.com/google/wire"
	"github.com/kalilventura/vehicle-management/internal/shared/domain/services"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	NewGooseMigrationService,
	wire.Bind(new(services.MigrationService), new(*GooseMigrationService)),
)

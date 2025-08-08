//go:build unit

package commands_test

import (
	"errors"
	"testing"

	shared "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/commands"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/test/vehicles/domain/builders"
	"github.com/kalilventura/vehicle-management/test/vehicles/infrastructure/repositories"
	"github.com/stretchr/testify/assert"
)

func TestListVehiclesCommand(t *testing.T) {
	t.Run("should call OnSuccess when there are vehicles to show", func(t *testing.T) {
		// given
		input := builders.NewListVehicleInputBuilder().Build()
		vehicles := builders.NewVehicleBuilder().BuildPagination()

		repository := repositories.NewInMemoryVehiclesRepository().WithVehicles(vehicles)
		command := commands.NewListVehiclesCommand(repository)

		// when & then
		listeners := commands.ListVehiclesListeners{
			OnSuccess: func(vehicles *shared.PaginatedEntity[entities.Vehicle]) {
				assert.NotNil(t, vehicles, "vehicles should not be nil")
			},
		}
		command.Execute(input, listeners)
	})

	t.Run("should call OnInternal when there's an error", func(t *testing.T) {
		// given
		input := builders.NewListVehicleInputBuilder().Build()
		repository := repositories.NewInMemoryVehiclesRepository().WithError(errors.New("list vehicle error"))
		command := commands.NewListVehiclesCommand(repository)

		// when & then
		listeners := commands.ListVehiclesListeners{
			OnInternalServerError: func(err error) {
				assert.NotNil(t, err, "err should not be nil")
			},
		}
		command.Execute(input, listeners)
	})
}

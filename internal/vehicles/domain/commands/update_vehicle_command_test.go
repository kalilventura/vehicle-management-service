//go:build unit

package commands_test

import (
	"errors"
	"testing"

	domainerror "github.com/kalilventura/vehicle-management/internal/shared/domain/errors"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/commands"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/test/vehicles/domain/builders"
	"github.com/kalilventura/vehicle-management/test/vehicles/infrastructure/repositories"
	"github.com/stretchr/testify/assert"
)

func TestUpdateVehicleCommand(t *testing.T) {
	t.Run("should call OnNotFound when there's no vehicle found", func(t *testing.T) {
		// given
		input := builders.NewUpdateVehicleInputBuilder().Build()
		repository := repositories.NewInMemoryVehiclesRepository().WithError(domainerror.ErrRecordNotFound)
		command := commands.NewUpdateVehicleCommand(repository)

		// when
		var wasCalled bool
		listeners := commands.UpdateVehicleListeners{
			OnNotFound: func() {
				wasCalled = true
			},
		}
		command.Execute(&input, listeners)

		// then
		assert.True(t, wasCalled, "OnNotFound was not called")
	})

	t.Run("should call OnInternalServerError when there's an error to find the vehicle", func(t *testing.T) {
		// given
		input := builders.NewUpdateVehicleInputBuilder().Build()
		repository := repositories.NewInMemoryVehiclesRepository().WithError(errors.New("get vehicle error"))
		command := commands.NewUpdateVehicleCommand(repository)

		// when & then
		listeners := commands.UpdateVehicleListeners{
			OnInternalServerError: func(err error) {
				assert.NotNil(t, err, "Error was not returned")
			},
		}
		command.Execute(&input, listeners)
	})

	t.Run("should call OnSuccess when the vehicle was updated", func(t *testing.T) {
		// given
		input := builders.NewUpdateVehicleInputBuilder().Build()
		vehicles := builders.NewVehicleBuilder().BuildPagination()
		repository := repositories.NewInMemoryVehiclesRepository().WithVehicles(vehicles)
		command := commands.NewUpdateVehicleCommand(repository)

		// when & then
		listeners := commands.UpdateVehicleListeners{
			OnSuccess: func(vehicle *entities.UpdateVehicleInput) {
				assert.NotNil(t, vehicle, "Vehicle was not returned")
			},
		}
		command.Execute(&input, listeners)
	})
}

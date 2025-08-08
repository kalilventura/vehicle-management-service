//go:build unit

package commands_test

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	domainerror "github.com/kalilventura/vehicle-management/internal/shared/domain/errors"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/commands"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/test/vehicles/domain/builders"
	"github.com/kalilventura/vehicle-management/test/vehicles/infrastructure/repositories"
	"github.com/stretchr/testify/assert"
)

func TestNewGetVehicleByIDCommand(t *testing.T) {
	t.Run("should call OnNotFound when there's no vehicle", func(t *testing.T) {
		// given
		repository := repositories.NewInMemoryVehiclesRepository().WithError(domainerror.ErrRecordNotFound)
		command := commands.NewGetVehicleByIDCommand(repository)

		// when
		var wasCalled bool
		listeners := commands.GetVehicleByIDListeners{
			OnNotFound: func() {
				wasCalled = true
			},
		}
		command.Execute(gofakeit.UUID(), listeners)

		// then
		assert.True(t, wasCalled, "OnNotFound was not called")
	})

	t.Run("should call OnInternalServerError when there's an error to find the vehicle", func(t *testing.T) {
		// given
		repository := repositories.NewInMemoryVehiclesRepository().WithError(errors.New("get vehicle error"))
		command := commands.NewGetVehicleByIDCommand(repository)

		// when & then
		listeners := commands.GetVehicleByIDListeners{
			OnInternalServerError: func(err error) {
				assert.NotNil(t, err, "Error was not returned")
			},
		}
		command.Execute(gofakeit.UUID(), listeners)
	})

	t.Run("should call OnSuccess when there's a vehicle to show", func(t *testing.T) {
		// given
		vehicles := builders.NewVehicleBuilder().BuildPagination()
		repository := repositories.NewInMemoryVehiclesRepository().WithVehicles(vehicles)
		command := commands.NewGetVehicleByIDCommand(repository)

		// when & then
		listeners := commands.GetVehicleByIDListeners{
			OnSuccess: func(vehicle *entities.Vehicle) {
				assert.NotNil(t, vehicle, "Vehicle was not returned")
			},
		}
		command.Execute(gofakeit.UUID(), listeners)
	})
}

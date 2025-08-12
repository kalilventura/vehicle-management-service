//go:build unit

package commands_test

import (
	"errors"
	"testing"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/commands"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/kalilventura/vehicle-management/test/vehicles/domain/builders"
	"github.com/kalilventura/vehicle-management/test/vehicles/infrastructure/repositories"
	"github.com/stretchr/testify/assert"
)

func TestSaveVehicleCommand(t *testing.T) {
	t.Run("should call OnNotValid when the vehicle is invalid", func(t *testing.T) {
		// given
		vehicle := builders.NewVehicleBuilder().BuildInvalid()
		repository := repositories.NewInMemoryVehiclesRepository()
		command := commands.NewSaveVehicleCommand(repository)

		// when & then
		listeners := commands.SaveVehicleListeners{
			OnNotValid: func(err error) {
				assert.Error(t, err)
			},
		}
		command.Execute(vehicle, listeners)
	})

	t.Run("should call OnInternalServerError when there's an error to save the new vehicle", func(t *testing.T) {
		// given
		mileage, _ := dtos.NewMileage(0)
		specification := entities.Specification{
			Mileage: mileage,
		}
		vehicle := builders.NewVehicleBuilder().
			WithSpecification(specification).
			WithCondition(dtos.New).
			Build()
		repository := repositories.NewInMemoryVehiclesRepository().WithError(errors.New("save vehicle error"))
		command := commands.NewSaveVehicleCommand(repository)

		// when & then
		listeners := commands.SaveVehicleListeners{
			OnInternalServerError: func(err error) {
				assert.Error(t, err)
			},
		}
		command.Execute(&vehicle, listeners)
	})

	t.Run("should call OnSuccess when the vehicle was saved successfully", func(t *testing.T) {
		// given
		mileage, _ := dtos.NewMileage(0)
		specification := entities.Specification{
			Mileage: mileage,
		}
		vehicle := builders.NewVehicleBuilder().
			WithSpecification(specification).
			WithCondition(dtos.New).
			Build()
		repository := repositories.NewInMemoryVehiclesRepository()
		command := commands.NewSaveVehicleCommand(repository)

		// when & then
		listeners := commands.SaveVehicleListeners{
			OnSuccess: func(vehicle *entities.Vehicle) {
				assert.NotNil(t, vehicle)
			},
		}
		command.Execute(&vehicle, listeners)
	})
}

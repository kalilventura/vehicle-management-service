//go:build unit

package commands_test

import (
  "errors"
  "testing"

  "github.com/kalilventura/vehicle-management/internal/vehicles/domain/commands"
  "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
  "github.com/kalilventura/vehicle-management/test/vehicles/domain/builders"
  "github.com/kalilventura/vehicle-management/test/vehicles/infrastructure/repositories"
  "github.com/kalilventura/vehicle-management/test/vehicles/infrastructure/services"
  "github.com/stretchr/testify/assert"
)

func TestSellVehicleCommand(t *testing.T) {
  t.Run("should call OnBadRequest when the application fails to pay", func(t *testing.T) {
    // given
    sellVehicle := builders.NewSellVehicleBuilder().Build()

    svc := services.NewPaymentsServiceStub().WithOnBadRequest()
    repo := repositories.NewInMemoryVehiclesRepository()
    command := commands.NewSellVehicleCommand(svc, repo)

    listeners := commands.SellVehicleListeners{
      OnBadRequest: func(err error) {
        assert.Error(t, err)
      },
    }

    // when
    command.Execute(&sellVehicle, listeners)
  })

  t.Run("should call OnInternalServerError due an unexpected fail", func(t *testing.T) {
    // given
    sellVehicle := builders.NewSellVehicleBuilder().Build()

    svc := services.NewPaymentsServiceStub().WithOnInternalServerError()
    repo := repositories.NewInMemoryVehiclesRepository()
    command := commands.NewSellVehicleCommand(svc, repo)

    listeners := commands.SellVehicleListeners{
      OnInternalServerError: func(err error) {
        assert.Error(t, err)
      },
    }

    // when
    command.Execute(&sellVehicle, listeners)
  })

  t.Run("should call OnInternalServerError due an unexpected error in the database", func(t *testing.T) {
    // given
    sellVehicle := builders.NewSellVehicleBuilder().Build()

    svc := services.NewPaymentsServiceStub().WithOnSuccess(&sellVehicle)
    repo := repositories.NewInMemoryVehiclesRepository().WithError(errors.New("save vehicle error"))
    command := commands.NewSellVehicleCommand(svc, repo)

    listeners := commands.SellVehicleListeners{
      OnInternalServerError: func(err error) {
        assert.Error(t, err)
      },
    }

    // when
    command.Execute(&sellVehicle, listeners)
  })

  t.Run("should call OnSuccess when the payment was saved", func(t *testing.T) {
    // given
    sellVehicle := builders.NewSellVehicleBuilder().Build()

    svc := services.NewPaymentsServiceStub().WithOnSuccess(&sellVehicle)
    repo := repositories.NewInMemoryVehiclesRepository()
    command := commands.NewSellVehicleCommand(svc, repo)

    listeners := commands.SellVehicleListeners{
      OnSuccess: func(sell *entities.SellVehicle) {
        assert.NotNil(t, sell)
      },
    }

    // when
    command.Execute(&sellVehicle, listeners)
  })
}

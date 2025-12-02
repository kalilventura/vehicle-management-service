//go:build integration

package persistence_test

import (
	"context"
	"testing"

	"github.com/kalilventura/vehicle-management/internal/shared/infrastructure/configuration"
	listvehicles "github.com/kalilventura/vehicle-management/internal/vehicles/application/use-cases/list-vehicles"
	entities2 "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	valueobjects "github.com/kalilventura/vehicle-management/internal/vehicles/domain/value-objects"
	persistence "github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/persistence"
	"github.com/kalilventura/vehicle-management/test/shared/infrastructure"
	"github.com/kalilventura/vehicle-management/test/vehicles/domain/builders"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormVehiclesRepositoryTestSuite struct {
	suite.Suite
	postgresContainer testcontainers.Container
	db                *gorm.DB
}

func (suite *GormVehiclesRepositoryTestSuite) SetupSuite() {
	ctx := context.Background()

	container, err := infrastructure.SetupPostgres(ctx)
	suite.Require().NoError(err)
	suite.postgresContainer = container

	settings, err := infrastructure.CreateDatabaseSettings(ctx, container)
	suite.Require().NoError(err)

	suite.db = configuration.NewDatabaseClient(settings)

	err = infrastructure.CreateDatabaseStructure(suite.db)
	suite.Require().NoError(err)
}

func (suite *GormVehiclesRepositoryTestSuite) TearDownSuite() {
	err := testcontainers.TerminateContainer(suite.postgresContainer)
	suite.Require().NoError(err)
	suite.T().Logf("Stopped postgres container")
}

func (suite *GormVehiclesRepositoryTestSuite) TestSuccessfully() {
	suite.Run("should create a new vehicle successfully", func() {
		// given
		mileage, _ := valueobjects.NewMileage(0)
		doors, _ := valueobjects.NewDoors(4)
		bodyType, _ := valueobjects.NewBodyType("sedan")
		transmission, _ := valueobjects.NewTransmission("automatic")
		fuelType, _ := valueobjects.NewFuelType("gasoline")

		specification := valueobjects.NewSpecification(valueobjects.SpecificationProps{
			Mileage:      mileage,
			Doors:        doors,
			BodyType:     bodyType,
			Transmission: transmission,
			FuelType:     fuelType,
			Engine:       "2.0L",
		})

		year, _ := valueobjects.NewYear(2006)
		vehicle := builders.NewVehicleBuilder().
			WithSpecification(specification).
			WithYear(year).
			Build()
		transaction := suite.db.Begin()
		defer transaction.Rollback()

		repository := persistence.NewGormVehiclesRepository(transaction)

		// when
		err := repository.Save(&vehicle)

		// then
		suite.NoError(err)
	})
}

func (suite *GormVehiclesRepositoryTestSuite) TestError() {
	suite.Run("should return an error when the application fails to save a vehicle", func() {
		// given
		vehicle := builders.NewVehicleBuilder().Build()
		transaction := suite.db.Begin()
		defer transaction.Rollback()

		repository := persistence.NewGormVehiclesRepository(transaction)

		// when
		err := repository.Save(&vehicle)

		// then
		suite.Error(err)
	})

	suite.Run("should return an error when the application fails to update a vehicle", func() {
		// given
		vehicle := builders.NewUpdateVehicleInputBuilder().Build()
		transaction := suite.db.Begin()
		defer transaction.Rollback()

		repository := persistence.NewGormVehiclesRepository(transaction)

		// when
		err := repository.Save(&vehicle)

		// then
		suite.Error(err)
	})

	suite.Run("should return an error when the application fails to get a vehicle", func() {
		// given
		transaction := suite.db.Begin()
		defer transaction.Rollback()

		repository := persistence.NewGormVehiclesRepository(transaction)

		// when
		_, err := repository.GetByID("")

		// then
		suite.Error(err)
	})

	suite.Run("should return an error when the application fails to list the vehicles", func() {
		// given
		input := listvehicles.Input{}
		dialector := postgres.Open("")
		transaction, _ := gorm.Open(dialector, &gorm.Config{})

		repository := persistence.NewGormVehiclesRepository(transaction)

		// when
		_, err := repository.FindWithFilters(input)

		// then
		suite.Error(err)
	})
}

func TestGormVehiclesRepository(t *testing.T) {
	suite.Run(t, new(GormVehiclesRepositoryTestSuite))
}

//go:build integration

package repositories_test

import (
	"context"
	"testing"

	"github.com/kalilventura/vehicle-management/internal/shared/infrastructure/configuration"
	entities2 "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/repositories"
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
		mileage, _ := dtos.NewMileage(0)
		doors, _ := dtos.NewDoors(4)
		specification := entities2.Specification{
			Mileage: mileage,
			Doors:   doors,
		}
		vehicle := builders.NewVehicleBuilder().
			WithSpecification(specification).
			WithYear(2006).
			Build()
		transaction := suite.db.Begin()
		defer transaction.Rollback()

		repository := repositories.NewGormVehiclesRepository(transaction)

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

		repository := repositories.NewGormVehiclesRepository(transaction)

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

		repository := repositories.NewGormVehiclesRepository(transaction)

		// when
		err := repository.Update(&vehicle)

		// then
		suite.Error(err)
	})

	suite.Run("should return an error when the application fails to get a vehicle", func() {
		// given
		transaction := suite.db.Begin()
		defer transaction.Rollback()

		repository := repositories.NewGormVehiclesRepository(transaction)

		// when
		_, err := repository.GetByID("")

		// then
		suite.Error(err)
	})

	suite.Run("should return an error when the application fails to list the vehicles", func() {
		// given
		input := dtos.ListVehiclesInput{}
		dialector := postgres.Open("")
		transaction, _ := gorm.Open(dialector, &gorm.Config{})

		repository := repositories.NewGormVehiclesRepository(transaction)

		// when
		_, err := repository.FindWithFilters(input)

		// then
		suite.Error(err)
	})
}

func TestGormVehiclesRepository(t *testing.T) {
	suite.Run(t, new(GormVehiclesRepositoryTestSuite))
}

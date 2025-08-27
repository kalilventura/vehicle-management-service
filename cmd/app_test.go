//go:build integration

package main_test

import (
	"context"
	"testing"

	main "github.com/kalilventura/vehicle-management/cmd"
	"github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/test/shared/infrastructure"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

type AppTestSuite struct {
	suite.Suite
	ctx               context.Context
	postgresContainer *postgres.PostgresContainer
	settings          *entities.DatabaseSettings
}

func (suite *AppTestSuite) SetupSuite() {
	ctx := context.Background()
	suite.ctx = ctx

	container, err := infrastructure.SetupPostgres(ctx)
	suite.Require().NoError(err)
	suite.postgresContainer = container

	settings, err := infrastructure.CreateDatabaseSettings(ctx, container)
	suite.Require().NoError(err)

	suite.settings = settings
}

func (suite *AppTestSuite) TearDownSuite() {
	err := testcontainers.TerminateContainer(suite.postgresContainer)
	suite.Require().NoError(err)
	suite.T().Logf("Stopped postgres container")
}

func (suite *AppTestSuite) TestAppSuccess() {
	suite.Run("should create the HTTP Modules successfully", func() {
		// given
		env, err := infrastructure.GetDatabaseEnvSettings(suite.ctx, suite.postgresContainer)
		suite.Require().NoError(err)

		err = infrastructure.SetEnvFromStruct(*env)
		suite.Require().NoError(err)

		// when
		modules := main.InjectModules()
		application := main.SetupServer(modules)

		// then
		suite.NotNil(application)

		err = infrastructure.UnsetEnvFromStruct(*env)
		suite.Require().NoError(err)
	})
}

func TestApp(t *testing.T) {
	suite.Run(t, new(AppTestSuite))
}

package infrastructure

import (
  "context"
  "net"

  "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
  "github.com/testcontainers/testcontainers-go/modules/postgres"
)

const (
  dbName     = "vehicle_management"
  dbUser     = "postgres"
  dbPassword = "postgres"
  img        = "postgres:18beta2-alpine3.22"
)

type SetupPostgresHelper struct {
  DbPort     string
  DbName     string
  DbUser     string
  DbPassword string
  DbHost     string
}

func SetupPostgres(ctx context.Context) (*postgres.PostgresContainer, error) {
  container, err := postgres.Run(ctx,
    img,
    postgres.WithDatabase(dbName),
    postgres.WithUsername(dbUser),
    postgres.WithPassword(dbPassword),
    postgres.BasicWaitStrategies(),
  )
  if err != nil {
    return nil, err
  }

  return container, nil
}

func GetDatabaseEnvSettings(ctx context.Context, container *postgres.PostgresContainer) (*SetupPostgresHelper, error) {
  host, port, err := getHostAndPort(ctx, container)
  if err != nil {
    return nil, err
  }

  setup := &SetupPostgresHelper{
    DbPort:     port,
    DbName:     dbName,
    DbUser:     dbUser,
    DbPassword: dbPassword,
    DbHost:     host,
  }
  return setup, nil
}

func CreateDatabaseSettings(ctx context.Context, container *postgres.PostgresContainer) (*entities.DatabaseSettings, error) {
  host, port, err := getHostAndPort(ctx, container)
  if err != nil {
    return nil, err
  }

  settings := entities.NewDatabaseSettings(host, dbName, port, dbUser, dbPassword, "disable")
  return settings, nil
}

func getHostAndPort(ctx context.Context, container *postgres.PostgresContainer) (string, string, error) {
  endpoint, err := container.PortEndpoint(ctx, "5432/tcp", "")
  if err != nil {
    return "", "", err
  }
  host, port, err := net.SplitHostPort(endpoint)
  if err != nil {
    return "", "", err
  }
  return host, port, nil
}

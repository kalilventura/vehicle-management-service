package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	_ "github.com/lib/pq" // required to migrate the database
	"github.com/pressly/goose/v3"
	"gorm.io/gorm"
)

type GooseMigrationService struct {
	gormDB   *gorm.DB
	settings *entities.DatabaseSettings
}

func NewGooseMigrationService(gormDB *gorm.DB, settings *entities.DatabaseSettings) *GooseMigrationService {
	return &GooseMigrationService{gormDB, settings}
}

func (s *GooseMigrationService) Run(databasePath string) error {
	db, sqlErr := sql.Open(s.gormDB.Name(), s.settings.GetDSN())
	if sqlErr != nil {
		return fmt.Errorf("failed to open database: %w", sqlErr)
	}
	defer db.Close()

	if err := goose.UpContext(context.Background(), db, databasePath); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	return nil
}

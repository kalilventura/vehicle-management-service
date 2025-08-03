package configuration

import (
	"github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newPostgresDialector(dsn string) gorm.Dialector {
	return postgres.Open(dsn)
}

func NewDatabaseClient(settings *entities.DatabaseSettings) *gorm.DB {
	dialector := newPostgresDialector(settings.GetDSN())
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		logger.WithField("error", err).
			Fatalf("🚨 Failed to connect to the database. Check your connection string and credentials.")
	}

	logger.WithFields(logger.Fields{
		"database_type": db.Dialector.Name(),
	}).Info("✅ Database connection established successfully.")
	return db
}

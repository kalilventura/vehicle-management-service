package services

type MigrationService interface {
	Run(databasePath string) error
}

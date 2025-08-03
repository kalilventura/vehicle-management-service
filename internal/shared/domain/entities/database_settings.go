package entities

import "fmt"

type DatabaseSettings struct {
	host     string
	name     string
	port     string
	user     string
	password string
	dbSSL    string
}

func NewDatabaseSettings(host, name, port, user, password, dbSSL string) *DatabaseSettings {
	return &DatabaseSettings{
		host,
		name,
		port,
		user,
		password,
		dbSSL,
	}
}

func (s DatabaseSettings) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		s.host, s.user, s.password, s.name, s.port, s.dbSSL,
	)
}

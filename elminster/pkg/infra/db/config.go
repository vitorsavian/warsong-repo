package db

import "github.com/jackc/pgx/v5"

type DatabaseConfig struct {
	Port int
	IP   string

	Login    string
	Password string

	Database string

	Connection *ConnectionClient
}

type ConnectionClient struct {
	Client *pgx.Conn
}

func CreateConfig() *DatabaseConfig {
	return &DatabaseConfig{}
}

package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/repository"
)

type DatabaseConfig struct {
	Port string
	IP   string

	Login    string
	Password string

	Database string
}

func CreateConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Login:    os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Port:     os.Getenv("DATABASE_PORT"),
		IP:       os.Getenv("DATABASE_IP"),
		Database: os.Getenv("DATABASE"),
	}
}

func (d *DatabaseConfig) CreateURL() string {
	return fmt.Sprintf("postgresql://%s:%s@%s/%s", d.Login, d.Password, d.IP, d.Database)
}

func (d *DatabaseConfig) CreatePoolConnection() (*repository.ConnectionClient, error) {
	connection := repository.ConnectionClient{}
	var err error

	connection.Client, err = pgx.Connect(context.Background(), d.CreateURL())
	if err != nil {
		return nil, err
	}

	return &connection, nil
}

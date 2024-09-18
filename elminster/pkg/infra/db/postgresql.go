package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (d *DatabaseConfig) CreateURL() string {
	return fmt.Sprintf("postgresql://%s:%s@%s/%s", d.Login, d.Password, d.IP, d.Database)
}

func (d *DatabaseConfig) CreatePoolConnection() error {
	connection := &ConnectionClient{}
	var err error

	connection.Client, err = pgx.Connect(context.Background(), d.CreateURL())
	if err != nil {
		return err
	}

	d.Connection = connection
	return nil
}

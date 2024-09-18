package db

import (
	"context"
	"github.com/jackz/pgx/v5"
	"os"
)

func CreatePoolConnection() {
	conn, err := pgx.Connect(context.Background(), "URL")
}

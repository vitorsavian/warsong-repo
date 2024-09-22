package repository

import (
	"github.com/jackc/pgx/v5"
)

type ConnectionClient struct {
	Client *pgx.Conn
}

func (c *ConnectionClient) CreateCharacter() {
	return

}

func (c *ConnectionClient) DeleteCharacter() {
	return
}

func (c *ConnectionClient) UpdateCharacter() {
	return
}

func (c *ConnectionClient) GetCharacter() {
	return
}

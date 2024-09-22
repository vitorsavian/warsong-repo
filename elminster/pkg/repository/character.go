package repository

import (
	"github.com/jackc/pgx/v5"
)

type ConnectionClient struct {
	Client *pgx.Conn
}

func (c *ConnectionClient) CreateCharacter() {

}

func (c *ConnectionClient) DeleteCharacter() {

}

func (c *ConnectionClient) UpdateCharacter() {

}

func (c *ConnectionClient) GetCharacter() {

}

package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/domain"
)

type ConnectionClient struct {
	Client *pgx.Conn
}

func (c *ConnectionClient) CreateCharacter(*domain.Character) error {

	return nil
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

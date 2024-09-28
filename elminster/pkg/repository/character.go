package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/domain"
)

type ConnectionClient struct {
	Client *pgx.Conn
}

var InsertCharacterSQL = `
  INSERT INTO CHARACTERS(id, name, level) VALUES($1, $2, $3)
`

func (c *ConnectionClient) CreateCharacter(char *domain.Character) error {
	c.Client.Exec(context.Background(), InsertCharacterSQL, char.Id, char.Name, char.Level)
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

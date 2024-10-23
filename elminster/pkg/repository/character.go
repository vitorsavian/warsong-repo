package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/domain"
)

type ConnectionClient struct {
	Client *pgx.Conn
}

var InsertCharacterSQL = `
	BEGIN TRANSACTION
   		INSERT INTO characters(id, name, level, id_stats) VALUES($1, $2, $3, $4)
   		INSERT INTO stats(id, hp, sp, str, dex, con, inte, wil, cha, sanity, courage) VALUES($4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15);
	COMMIT;
`

func (c *ConnectionClient) CreateCharacter(char *domain.Character) error {
	if _, err := c.Client.Query(context.Background(), InsertCharacterSQL, char.Id, char.Name, char.Level,
		char.Id, char.HP, char.SP, char.Str,
		char.Dex, char.Con, char.Int, char.Wil,
		char.Cha, char.Sanity, char.Courage,
	); err != nil {
		logrus.Errorf("Failed to exec query for create character: %v", err)
		return err
	}

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

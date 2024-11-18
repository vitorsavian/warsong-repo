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

var InsertStatsCharacterSQL = `
  INSERT INTO stats(id, hp, sp, str, dex, con, inte, wil, cha, sanity, courage)
  VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
`

var InsertCharacterSQL = `
  INSERT INTO characters(id, name, level, id_stats)
  VALUES($1, $2, $3, $4);
`

var GetCharacterSQL = `
  SELECT characters.id, characters.name, characters.level, stats.hp, stats.sp, stats.str, stats.dex, stats.con, stats.inte, stats.wil, stats.cha, stats.sanity, stats.courage
  FROM characters INNER JOIN stats ON characters.id_stats = stats.id WHERE characters.id = $1;
`

var DeleteStatsCharacterSQL = `
  DELETE FROM stats WHERE id = $1;
`

var DeleteCharacterSQL = `
  DELETE FROM characters WHERE id = $1;
`

var UpdateCharacterSQL = `
`

func (c *ConnectionClient) CreateCharacter(char *domain.Character) error {
	ctx := context.Background()

	tx, err := c.Client.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, InsertStatsCharacterSQL,
		char.Id, char.HP, char.SP, char.Str,
		char.Dex, char.Con, char.Int, char.Wil,
		char.Cha, char.Sanity, char.Courage,
	); err != nil {
		logrus.Errorf("Unable to exec query for create stats for a character: %v", err)
		return err
	}

	if _, err := tx.Exec(context.Background(), InsertCharacterSQL,
		char.Id, char.Name, char.Level, char.Id,
	); err != nil {
		logrus.Errorf("Unable to exec query for create character: %v", err)
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (c *ConnectionClient) DeleteCharacter(id string) error {
	ctx := context.Background()

	tx, err := c.Client.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, DeleteStatsCharacterSQL, id); err != nil {
		logrus.Errorf("Unable to exec query for delete stats for a character: %v", err)
		return err
	}

	if _, err := tx.Exec(context.Background(), DeleteCharacterSQL, id); err != nil {
		logrus.Errorf("Unable to exec query for delete character: %v", err)
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (c *ConnectionClient) UpdateCharacter(char *domain.Character) error {
	return nil
}

func (c *ConnectionClient) GetCharacter(id string) (*domain.Character, error) {
	result, err := c.Client.Query(context.Background(), GetCharacterSQL, id)
	if err != nil {
		logrus.Errorf("Unable to get character: %v", err)
		return nil, err
	}

	defer result.Close()

	var character domain.Character
	for result.Next() {
		err := result.Scan(&character.Id, &character.Name, &character.Level,
			&character.HP, &character.SP, &character.Str, &character.Dex, &character.Con,
			&character.Int, &character.Wil, &character.Cha, &character.Sanity, &character.Courage)

		if err != nil {
			logrus.Errorf("Unable to get rows when getting a character: %v", err)
			return nil, err
		}
	}

	return &character, nil
}

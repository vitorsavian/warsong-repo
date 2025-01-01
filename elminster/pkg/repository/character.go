package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/domain"
)

var InsertStatsCharacterSQL = `
  	INSERT INTO cstats(id, hp, sp, str, dex, con, inte, wil, cha, sanity, courage)
  	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
`

var InsertCharacterSQL = `
  	INSERT INTO characters(id, name, level, id_stats)
  	VALUES($1, $2, $3, $4);
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

var DeleteStatsCharacterSQL = `
  	DELETE FROM cstats WHERE id = $1;
`

var DeleteCharacterSQL = `
	DELETE FROM characters WHERE id = $1;
`

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

var UpdateCharacterSQL = `
	UPDATE characters
	SET name = $1, level = $2,
	WHERE id = $3;
`

var UpdateStatsSQL = `
	UPDATE cstats
	SET hp = $1, sp = $2, str = $3, dex = $4, con = $5, inte = $6, wil = $7, cha = $8, sanity = $9, courage = $10
	WHERE id = $11;
`

func (c *ConnectionClient) UpdateCharacter(char *domain.Character) error {
	ctx := context.Background()

	tx, err := c.Client.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, UpdateCharacterSQL, char.Name, char.Level, char.Id); err != nil {
		logrus.Errorf("Unable to exec query for update character: %v", err)
		return err
	}

	if _, err := tx.Exec(ctx, UpdateStatsSQL,
		char.HP, char.SP, char.Str,
		char.Dex, char.Con, char.Int,
		char.Wil, char.Cha, char.Sanity,
		char.Courage, char.Id); err != nil {
		logrus.Errorf("Unable to exec query for update stats: %v", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

var GetCharacterSQL = `
  	SELECT characters.id, characters.name, characters.level, cstats.hp, cstats.sp, cstats.str, cstats.dex, cstats.con, cstats.inte, cstats.wil, cstats.cha, cstats.sanity, cstats.courage
  	FROM characters INNER JOIN cstats ON characters.id_stats = stats.id WHERE characters.id = $1;
`

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

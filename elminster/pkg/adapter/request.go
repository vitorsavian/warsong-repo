package adapter

import "errors"

type CharacterCreationRequestAdapter struct {
	Level int8   `json:"level"`
	Name  string `json:"name"`

	Stats StatsCharacterCreationRequestAdapter `json:"stats"`
}

type StatsCharacterCreationRequestAdapter struct {
	Str int8 `json:"str"`
	Dex int8 `json:"dex"`
	Con int8 `json:"con"`
	Cha int8 `json:"cha"`
	Int int8 `json:"int"`
	Wil int8 `json:"wil"`
}

func CheckCreateCharacterBody(body *CharacterCreationRequestAdapter) error {
	if body.Level <= 0 {
		return errors.New("level is below zero or zero")
	}

	if body.Name == "" {
		return errors.New("name is blank")
	}

	if body.Stats.Str < 0 || body.Stats.Dex < 0 || body.Stats.Con < 0 || body.Stats.Cha < 0 || body.Stats.Wil < 0 || body.Stats.Int < 0 {
		return errors.New("stats for the character are negative")
	}

	return nil
}

type CharacterUpdateRequestAdapter struct {
}

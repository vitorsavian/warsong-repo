package domain

import (
	"errors"

	"github.com/google/uuid"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/adapter"
)

type Character struct {
	Id    string
	Level int8
	Name  string

	HP        int16
	CurrentHP int16

	SP        int16
	CurrentSP int16

	Str int8
	Dex int8
	Con int8
	Cha int8
	Int int8
	Wil int8

	Sanity  int16
	Courage int16

	Feat []Feat
}

func ValidateCharacter(body *Character) error {
	if body.Level <= 0 {
		return errors.New("level is below zero or zero")
	}

	if body.Name == "" {
		return errors.New("name is blank")
	}

	if body.Str < 0 || body.Dex < 0 || body.Con < 0 || body.Cha < 0 || body.Wil < 0 || body.Int < 0 {
		return errors.New("stats for the character are negative")
	}

	return nil
}

func CreateCharacter(body *adapter.CharacterCreationRequestAdapter) *Character {
	newCharacter := &Character{
		Id:    uuid.New().String(),
		Level: body.Level,
		Name:  body.Name,
		Str:   body.Stats.Str,
		Dex:   body.Stats.Dex,
		Con:   body.Stats.Con,
		Cha:   body.Stats.Cha,
		Int:   body.Stats.Int,
		Wil:   body.Stats.Wil,
	}

	newCharacter.SetHP()
	newCharacter.SetSP()
	newCharacter.SetSanity()
	newCharacter.SetCourage()

	return newCharacter
}

func UpdateCharacter(body *adapter.CharacterUpdateRequestAdapter) *Character {
	character := &Character{
		Id:    body.Id,
		Level: body.Level,
		Name:  body.Name,
		Str:   body.Stats.Str,
		Dex:   body.Stats.Dex,
		Con:   body.Stats.Con,
		Cha:   body.Stats.Cha,
		Int:   body.Stats.Int,
		Wil:   body.Stats.Wil,
	}

	character.SetHP()
	character.SetSP()
	character.SetCourage()
	character.SetSanity()

	return character
}

func (c *Character) LevelUp() {
	c.Level++
	c.HP += int16(c.Con)
	c.SP += int16(c.Int)
}

func (c *Character) Rest() {
	c.CurrentHP = c.HP
	c.CurrentSP = c.SP
}

func (c *Character) SetHP() {
	c.HP = int16(c.Level * c.Con)
	c.CurrentHP = c.HP
}

func (c *Character) SetSP() {
	c.SP = int16(c.Level * c.Int)
	c.CurrentSP = c.SP
}

func (c *Character) SetSanity() {
	c.Sanity = 45 + int16(c.Level+c.Wil-(c.Int/2))
}

func (c *Character) SetCourage() {
	c.Courage = 30 + int16(c.Level+c.Wil)
}

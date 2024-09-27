package domain

import "github.com/vitorsavian/warsong-repo/elminster/pkg/adapter"

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

type Feat struct {
	Name        string
	Description string
}

func CreateCharacter(body *adapter.CharacterCreationRequestAdapter) (*Character, error) {
	newCharacter := &Character{
		Id:    "test",
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

	return newCharacter, nil
}

func (c *Character) LevelUp() {
	c.Level++
	c.HP += int16(c.Con)
	c.SP += int16(c.Int)
}

func (c *Character) Rest() {

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

package character

type Character struct {
	Id   string
	Name string
	HP   int16
	SP   int16

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

func CreateCharacter() *Character {

	return nil
}

func (c *Character) LevelUp() {
	c.HP += int16(c.Con)
	c.SP += int16(c.Int)

}

func (c *Character) Fight(enemy *Character) {

}

func (c *Character) Rest() {

}

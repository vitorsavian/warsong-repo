package request

type CharacterCreationRequest struct {
	Level int8
	Name  string

	Str int8
	Dex int8
	Con int8
	Cha int8
	Int int8
	Wil int8
}

type CharacterUpdateRequest struct {
}

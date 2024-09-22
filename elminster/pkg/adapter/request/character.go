package request

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

type CharacterUpdateRequestAdapter struct {
}

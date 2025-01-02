package adapter

type CharacterCreationRequestAdapter struct {
	Id    string `json:"id"`
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
	Id    string
	Level int8   `json:"level"`
	Name  string `json:"name"`

	Stats StatsCharacterUpdateRequestAdapter `json:"stats"`
}

type StatsCharacterUpdateRequestAdapter struct {
	Str int8 `json:"str"`
	Dex int8 `json:"dex"`
	Con int8 `json:"con"`
	Cha int8 `json:"cha"`
	Int int8 `json:"int"`
	Wil int8 `json:"wil"`
}

type GenericResponseAdapter struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

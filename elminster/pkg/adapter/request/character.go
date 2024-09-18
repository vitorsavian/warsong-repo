package request

import "github.com/labstack/echo/v4"

type CharacterCreationRequest struct {
	Level int8   `json:"level"`
	Name  string `json:"name"`

	Stats StatsCharacterCreationRequest `json:"stats"`
}

type StatsCharacterCreationRequest struct {
	Str int8 `json:"str"`
	Dex int8 `json:"dex"`
	Con int8 `json:"con"`
	Cha int8 `json:"cha"`
	Int int8 `json:"int"`
	Wil int8 `json:"wil"`
}

type CharacterUpdateRequest struct {
}

func CreateCharacter(c echo.Context) error {
	return nil
}

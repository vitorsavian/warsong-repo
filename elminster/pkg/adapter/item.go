package adapter

type ItemCreationRequestAdapter struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Rarity int8   `json:"level"`
}

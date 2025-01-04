package domain

type ItemType string

type Item struct {
	Value int32

	Type ItemType

	IsEquipped bool
}

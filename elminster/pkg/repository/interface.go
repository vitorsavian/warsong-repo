package repository

import "github.com/vitorsavian/warsong-repo/elminster/pkg/domain"

type ICharacter interface {
	CreateCharacter(*domain.Character) error
	DeleteCharacter(id string) error
	UpdateCharacter(*domain.Character) error
	GetCharacter(id string) (*domain.Character, error)
}

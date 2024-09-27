package repository

import "github.com/vitorsavian/warsong-repo/elminster/pkg/domain"

type ICharacter interface {
	CreateCharacter(*domain.Character) error
	DeleteCharacter()
	UpdateCharacter()
	GetCharacter()
}

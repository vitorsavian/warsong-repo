package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/domain"
)

type ConnectionClient struct {
	Client *pgx.Conn
}

type ICharacter interface {
	CreateCharacter(*domain.Character) error
	DeleteCharacter(id string) error
	UpdateCharacter(*domain.Character) error
	GetCharacter(id string) (*domain.Character, error)
}

type IFeat interface {
	CreateFeat(*domain.Feat) error
	DeleteFeat(id string) error
	UpdateFeat(*domain.Feat) error
	GetFeat(id string) (*domain.Feat, error)
}

type IItem interface {
	CreateItem() error
	DeleteItem() error
	UpdateItem() error
	GetItem() error
}

type IMonster interface {
	CreateMonster() error
	DeleteMonster() error
	UpdateMonster() error
	GetMonster() error
}

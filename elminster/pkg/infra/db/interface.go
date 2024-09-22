package db

import "github.com/vitorsavian/warsong-repo/elminster/pkg/repository"

type IDatabase interface {
	CreateURL() string
	CreatePoolConnection() (*repository.ICharacter, error)
}

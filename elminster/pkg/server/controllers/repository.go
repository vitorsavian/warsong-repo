package controllers

import "github.com/vitorsavian/warsong-repo/elminster/pkg/repository"

type RepositoryController struct {
	CharacterRepo *repository.Character
}

func CreateRepo() *RepositoryController {
	return nil
}

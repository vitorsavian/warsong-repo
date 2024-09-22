package controllers

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/infra/db"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/repository"
)

var lockCharacterController = &sync.Mutex{}

type CharacterController struct {
	CharacterRepo *repository.ICharacter
}

var CharacterControllerStance *CharacterController

func GetController() *CharacterController {
	if CharacterControllerStance == nil {
		lockCharacterController.Lock()
		defer lockCharacterController.Unlock()

		if CharacterControllerStance == nil {
			databaseConfig := db.CreateConfig()
			repo, err := databaseConfig.CreatePoolConnection()
			if err != nil {
				logrus.Panic("Panic for test")
			}

			CharacterControllerStance = &CharacterController{
				CharacterRepo: repo,
			}
		}
	}

	return CharacterControllerStance
}

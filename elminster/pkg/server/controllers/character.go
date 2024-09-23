package controllers

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/adapter"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/infra/db"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/repository"
)

var lockCharacterController = &sync.Mutex{}

type CharacterController struct {
	CharacterRepo repository.ICharacter
}

var CharacterControllerStance *CharacterController

func GetController() (*CharacterController, error) {
	if CharacterControllerStance == nil {
		logrus.Info("Creating controller")

		lockCharacterController.Lock()
		defer lockCharacterController.Unlock()

		if CharacterControllerStance == nil {
			databaseConfig := db.CreateConfig()
			repo, err := databaseConfig.CreatePoolConnection()
			if err != nil {
				logrus.Errorf("Error while creating the pool connection: %s", err.Error())
				return nil, err
			}

			CharacterControllerStance = &CharacterController{
				CharacterRepo: repo,
			}
		}
	}

	logrus.Info("Controller delivered")
	return CharacterControllerStance, nil
}

func (c *CharacterController) CreateCharacter(body *adapter.CharacterCreationRequestAdapter) error {
	return nil
}

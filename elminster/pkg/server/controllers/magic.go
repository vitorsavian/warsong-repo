package controllers

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/infra/db"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/repository"
)

type MagicController struct {
	MagicRepo repository.ICharacter
}

var MagicControllerStance *MagicController

var lockMagicController = &sync.Mutex{}

func getMagicController() (*MagicController, error) {
	if MagicControllerStance == nil {
		logrus.Info("Creating magic controller")

		lockMagicController.Lock()
		defer lockItemController.Unlock()

		if MagicControllerStance == nil {
			databaseConfig := db.CreateConfig()
			repo, err := databaseConfig.CreatePoolConnection()
			if err != nil {
				logrus.Errorf("Unable to create pool connection for item controller: %v", err)
				return nil, err
			}

			MagicControllerStance = &MagicController{
				MagicRepo: repo,
			}
		}
	}

	logrus.Info("Magic controller delivered")
	return MagicControllerStance, nil
}

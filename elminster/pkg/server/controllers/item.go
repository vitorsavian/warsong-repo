package controllers

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/infra/db"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/repository"
)

type ItemController struct {
	ItemRepo repository.ICharacter
}

var ItemControllerStance *ItemController

var lockItemController = &sync.Mutex{}

func getItemController() (*ItemController, error) {
	if ItemControllerStance == nil {
		logrus.Info("Creating item controller")

		lockItemController.Lock()
		defer lockItemController.Unlock()

		if ItemControllerStance == nil {
			databaseConfig := db.CreateConfig()
			repo, err := databaseConfig.CreatePoolConnection()
			if err != nil {
				logrus.Errorf("Unable to create pool connection for item controller: %v", err)
				return nil, err
			}

			ItemControllerStance = &ItemController{
				ItemRepo: repo,
			}
		}
	}

	logrus.Info("Item controller delivered")
	return ItemControllerStance, nil
}

package controllers

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/infra/db"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/repository"
)

type MonsterController struct {
	MonsterRepo repository.ICharacter
}

var MonsterControllerStance *MonsterController

var lockMonsterController = &sync.Mutex{}

func getMonsterController() (*MonsterController, error) {
	if MonsterControllerStance == nil {
		logrus.Info("Creating item controller")

		lockMonsterController.Lock()
		defer lockMonsterController.Unlock()

		if MonsterControllerStance == nil {
			databaseConfig := db.CreateConfig()
			repo, err := databaseConfig.CreatePoolConnection()
			if err != nil {
				logrus.Errorf("Unable to create pool connection for monster controller: %v", err)
				return nil, err
			}

			MonsterControllerStance = &MonsterController{
				MonsterRepo: repo,
			}
		}
	}

	logrus.Info("Monster controller delivered")
	return MonsterControllerStance, nil
}

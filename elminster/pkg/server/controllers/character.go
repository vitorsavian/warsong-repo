package controllers

import (
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/adapter"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/domain"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/infra/db"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/repository"
)

var lockCharacterController = &sync.Mutex{}

type CharacterController struct {
	CharacterRepo repository.ICharacter
}

var CharacterControllerStance *CharacterController

func GetCharacterController() (*CharacterController, error) {
	if CharacterControllerStance == nil {
		logrus.Info("Creating controller")

		lockCharacterController.Lock()
		defer lockCharacterController.Unlock()

		if CharacterControllerStance == nil {
			databaseConfig := db.CreateConfig()
			repo, err := databaseConfig.CreatePoolConnection()
			if err != nil {
				logrus.Errorf("Unable to create pool connection: %v", err)
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

func (c *CharacterController) CreateCharacter(body *adapter.CharacterCreationRequestAdapter) (int, error) {
	if err := adapter.CheckCreateBody(body); err != nil {
		return http.StatusBadRequest, err
	}

	char, err := domain.CreateCharacter(body)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = c.CharacterRepo.CreateCharacter(char)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (c *CharacterController) UpdateCharacter(body *adapter.CharacterUpdateRequestAdapter) (int, error) {
	if err := adapter.CheckUpdateBody(body); err != nil {
		return http.StatusBadRequest, err
	}

	char, err := domain.UpdateCharacter(body)
	if err != nil {
		return http.StatusBadRequest, err
	}

	if err := c.CharacterRepo.UpdateCharacter(char); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}

func (c *CharacterController) DeleteCharacter(id string) (int, error) {
	if err := c.CharacterRepo.DeleteCharacter(id); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}

func (c *CharacterController) GetCharacter(id string) (*domain.Character, int, error) {
	character, err := c.CharacterRepo.GetCharacter(id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return character, http.StatusFound, nil
}

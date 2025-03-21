package tcp

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/adapter"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/server/controllers"
)

func CreateCharacter(c echo.Context) error {
	body := new(adapter.CharacterCreationRequestAdapter)
	if err := c.Bind(body); err != nil {
		logrus.Errorf("Error found in Create Character body bind: %s", err.Error())
		c.JSON(http.StatusInternalServerError, &adapter.GenericResponseAdapter{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return err
	}

	controller, err := controllers.GetCharacterController()
	if err != nil {
		logrus.Errorf("Unable to get character controller: %s", err.Error())
		c.JSON(http.StatusInternalServerError, &adapter.GenericResponseAdapter{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})

		return nil
	}

	status, err := controller.CreateCharacter(body)
	if err != nil {
		c.JSON(status, &adapter.GenericResponseAdapter{
			Status:  status,
			Message: err.Error(),
		})

		return nil
	}

	c.JSON(status, &adapter.GenericResponseAdapter{
		Status:  status,
		Message: "Character created",
	})
	return nil
}

func GetCharacter(c echo.Context) error {
	controller, err := controllers.GetCharacterController()
	if err != nil {
		logrus.Errorf("Unable to get character controller: %s", err.Error())
		c.JSON(http.StatusInternalServerError, &adapter.GenericResponseAdapter{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})

		return nil
	}

	id := c.Param("id")

	character, status, err := controller.GetCharacter(id)
	if err != nil {
		c.JSON(status, &adapter.GenericResponseAdapter{
			Status:  status,
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusCreated, character)
	return nil
}

func DeleteCharacter(c echo.Context) error {
	controller, err := controllers.GetCharacterController()
	if err != nil {
		logrus.Errorf("Unable to get character controller: %s", err.Error())
		c.JSON(http.StatusInternalServerError, &adapter.GenericResponseAdapter{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})

		return nil
	}

	id := c.Param("id")

	status, err := controller.DeleteCharacter(id)
	if err != nil {
		c.JSON(status, &adapter.GenericResponseAdapter{
			Status:  status,
			Message: err.Error(),
		})
		return nil
	}

	c.JSON(status, &adapter.GenericResponseAdapter{
		Status:  status,
		Message: "Characte  deleted",
	})

	return nil
}

func UpdateCharacter(c echo.Context) error {
	body := new(adapter.CharacterUpdateRequestAdapter)
	if err := c.Bind(body); err != nil {
		logrus.Errorf("Unable to bind body in Update Character: %s", err.Error())
		c.JSON(http.StatusInternalServerError, &adapter.GenericResponseAdapter{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})

		return err
	}

	body.Id = c.Param("id")

	controller, err := controllers.GetCharacterController()
	if err != nil {
		logrus.Errorf("Unable to get character controller: %s", err.Error())
		c.JSON(http.StatusInternalServerError, &adapter.GenericResponseAdapter{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})

		return nil
	}

	status, err := controller.UpdateCharacter(body)
	if err != nil {
		c.JSON(status, &adapter.GenericResponseAdapter{
			Status:  status,
			Message: err.Error(),
		})
		return nil
	}

	c.JSON(status, &adapter.GenericResponseAdapter{
		Status:  status,
		Message: "Character updated",
	})

	return nil
}

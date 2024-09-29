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
		return err
	}

	controller, err := controllers.GetController()
	if err != nil {
		logrus.Panic()
		return nil
	}

	status, err := controller.CreateCharacter(body)
	if err != nil {
		c.JSON(status, "")
		return nil
	}

	c.JSON(http.StatusCreated, "")
	return nil
}

func GetCharacter(c echo.Context) error {
	body := new(adapter.CharacterCreationRequestAdapter)
	if err := c.Bind(body); err != nil {
		logrus.Errorf("Error found in Create Character body bind: %s", err.Error())
		return err
	}

	controller, err := controllers.GetController()
	if err != nil {
		logrus.Panic()
		return nil
	}

	status, err := controller.CreateCharacter(body)
	if err != nil {
		c.JSON(status, "")
		return nil
	}

	c.JSON(http.StatusCreated, "")
	return nil
}

func DeleteCharacter(c echo.Context) error {
	body := new(adapter.CharacterCreationRequestAdapter)
	if err := c.Bind(body); err != nil {
		logrus.Errorf("Error found in Create Character body bind: %s", err.Error())
		return err
	}

	controller, err := controllers.GetController()
	if err != nil {
		logrus.Panic()
		return nil
	}

	status, err := controller.CreateCharacter(body)
	if err != nil {
		c.JSON(status, "")
		return nil
	}

	c.JSON(http.StatusCreated, "")
	return nil
}

func UpdateCharacter(c echo.Context) error {
	body := new(adapter.CharacterCreationRequestAdapter)
	if err := c.Bind(body); err != nil {
		logrus.Errorf("Error found in Create Character body bind: %s", err.Error())
		return err
	}

	controller, err := controllers.GetController()
	if err != nil {
		logrus.Panic()
		return nil
	}

	status, err := controller.CreateCharacter(body)
	if err != nil {
		c.JSON(status, "")
		return nil
	}

	c.JSON(http.StatusCreated, "")
	return nil
}

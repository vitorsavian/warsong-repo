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

	if err = controller.CreateCharacter(body); err != nil {
		return nil
	}

	c.JSON(http.StatusCreated, "")
	return nil
}

package tcp

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/adapter"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/server/controllers"
)

func CreateCharacter(c echo.Context) error {
	logrus.Panic("testing")
	body := new(adapter.CharacterCreationRequestAdapter)
	if err := c.Bind(body); err != nil {
		return err
	}

	logrus.Panic(body)

	controller, err := controllers.GetController()
	if err != nil {
		return nil
	}

	if err = controller.CreateCharacter(body); err != nil {
		return nil
	}

	c.JSON(http.StatusCreated, "")
	return nil
}

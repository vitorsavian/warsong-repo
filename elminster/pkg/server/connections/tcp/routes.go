package tcp

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (e *Config) SetupRoutes() {
	api := e.Framework.Conn.Group("/api")

	api.GET("/healthz", Healthz)

	character := api.Group("/character")
	character.GET("/:id", GetCharacter)
	character.PUT("/:id", UpdateCharacter)
	character.DELETE("/:id", DeleteCharacter)
	character.POST("", CreateCharacter)

	_ = api.Group("/feat")

	_ = api.Group("/weapon")
}

func Healthz(c echo.Context) error {
	logrus.Println("Healthz called")
	c.NoContent(http.StatusOK)
	return nil
}

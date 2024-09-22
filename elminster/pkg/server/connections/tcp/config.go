package tcp

import (
	"log"

	"github.com/labstack/echo/v4"
)

func ConfigServer() (*Config, error) {
	createConfig := &Config{
		IPv4:      "",
		Port:      3000,
		Framework: &EchoConfig{},
	}

	createConfig.Framework.CreateServer()
	return createConfig, nil
}

func (c *EchoConfig) CreateServer() {
	e := echo.New()

	c.Conn = e
}

func (c *EchoConfig) Listen() {
	c.SetupRoutes()

	if err := c.Conn.Server.ListenAndServe(); err != nil {
		log.Fatal("Fatal")
	}
}

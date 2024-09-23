package tcp

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func ConfigServer() (*Config, error) {
	createConfig := &Config{
		IPv4:      "",
		Port:      os.Getenv("SERVER_TCP_PORT"),
		Framework: &EchoConfig{},
	}

	createConfig.Framework.CreateServer()
	return createConfig, nil
}

func (c *EchoConfig) CreateServer() {
	e := echo.New()

	c.Conn = e
}

func (c *Config) Listen() {
	c.SetupRoutes()

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", c.IPv4, c.Port),
		Handler: c.Framework.Conn,
	}

	logrus.Infof("listen server in: %s:%s", c.IPv4, c.Port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logrus.Fatal(err)
	}
}

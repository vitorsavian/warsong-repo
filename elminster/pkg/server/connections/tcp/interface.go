package tcp

import "github.com/labstack/echo/v4"

type Config struct {
	IPv4      string
	Port      int
	Framework *EchoConfig
}

type EchoConfig struct {
	Conn *echo.Echo
}

type Connection interface {
	CreateServer()
	Listen()
}
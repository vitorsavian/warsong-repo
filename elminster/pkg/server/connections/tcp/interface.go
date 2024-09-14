package tcp

import "github.com/labstack/echo/v4"

type EchoConfig struct {
	IPv4 string
	Port int
	Conn *echo.Echo
}

type Connection interface {
	CreateServer()
	Listen()
}

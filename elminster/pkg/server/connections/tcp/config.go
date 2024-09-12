package tcp

import "github.com/labstack/echo/v4"

type Server struct {
	IPv4 string
	Port int
}

func CreateServer(config Server) *echo.Echo {
	e := echo.New()

	return e
}

func Listen(server *echo.Echo) {

}

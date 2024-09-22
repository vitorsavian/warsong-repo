package tcp

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func CreateCharacter(c echo.Context) error {
	fmt.Println("testing this")
	return nil
}

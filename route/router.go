package route

import (
	"echo-gorm/api"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", api.Home)

	e.GET("/users", api.GetUsers)
	e.POST("/users", api.PostUsers)

	return e
}

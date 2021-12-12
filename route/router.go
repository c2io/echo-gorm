package route

import (
	"echo-gorm/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())    // log request info
	e.Use(middleware.Recover())   // auto recover from any panic
	e.Use(middleware.RequestID()) // log request info with id

	e.GET("/", api.Home)

	e.GET("/users", api.GetUsers)
	e.POST("/users", api.PostUsers)

	return e
}

package api

import (
	"echo-gorm/db"
	"echo-gorm/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	db := db.DbManager()
	users := []model.User{}
	db.Find(&users)
	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)
	return c.JSON(http.StatusOK, users)
}

func PostUsers(c echo.Context) error {
	user := new(model.User)
	c.Bind(user)

	db := db.DbManager()
	db.Create(user)

	return c.NoContent(http.StatusCreated)
}

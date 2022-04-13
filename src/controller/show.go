package controller

import (
	"go-react/database"
	"go-react/model"
	"net/http"

	"github.com/labstack/echo"
)

// ハンドラーを定義
func UserTest() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(model.User)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusOK, nil)
		}
		return c.JSON(http.StatusOK, u)
	}
}
func ShowUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := database.Open()
		user := model.User{}
		id := c.Param("id")
		db.Where("id = ?", id).First(&user)
		if user.ID == 0 {
			return c.JSON(http.StatusOK, nil)
		}
		return c.JSON(http.StatusOK, user)
	}
}
func CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := database.Open()
		user := new(model.User)
		if err := c.Bind(user); err != nil {
			return err
		}

		// user := model.User{Name: name, Email: email}

		result := db.Create(&user) // pass pointer of data to Create

		return c.JSON(http.StatusOK, result.Error)

	}
}

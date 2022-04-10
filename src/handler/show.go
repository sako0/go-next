package handler

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
		db := database.Connect()
		user := model.User{}
		id := c.Param("id")
		db.Where("id = ?", id).First(&user)
		if user.ID == "" {
			return c.JSON(http.StatusOK, nil)
		}
		return c.JSON(http.StatusOK, user)
	}
}

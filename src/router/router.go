package router

import (
	"go-react/controller"

	"github.com/labstack/echo"
)

func Handler() *echo.Echo {
	// echoのインスタンスを作成
	e := echo.New()
	// ルートを設定
	e.GET("/user/:id", controller.ShowUser())
	e.POST("/user", controller.CreateUser())
	return e
}

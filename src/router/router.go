package router

import (
	"go-react/handler"

	"github.com/labstack/echo"
)

func Handler() *echo.Echo {
	// echoのインスタンスを作成
	e := echo.New()
	// ルートを設定
	e.GET("/user/:id", handler.ShowUser())
	return e
}

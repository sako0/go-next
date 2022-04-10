package main

import "go-react/router"

func main() {
	router := router.Handler()
	// サーバーをポート番号1323で起動
	router.Logger.Fatal(router.Start(":1323"))
}

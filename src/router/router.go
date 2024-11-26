package router

import (
	// 振り分ける API
	"restapi/modules/testapi"

	"github.com/labstack/echo/v4"
)

/**
 * API のルーティングを管理する
 */
func Routing(e *echo.Echo) {
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.GET("/testapi/all", testapi.All)       // 全件取得
	e.GET("/testapi/:id", testapi.Content)   // ID を指定して取得
	e.POST("/testapi", testapi.Register)     // 登録
	e.PUT("/testapi/:id", testapi.Update)    // ID を指定して更新
	e.DELETE("/testapi/:id", testapi.Delete) // ID を指定して削除
}

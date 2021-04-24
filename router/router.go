/*
	Created by XiaoInk at 2021/04/23 16:23:17
	GitHub: https://github.com/XiaoInk
*/

package router

import (
	"net/http"

	"github.com/XiaoInk/GPL/handler"
	"github.com/gin-gonic/gin"
)

var App *gin.Engine

func init() {
	App = gin.Default()

	// 可作健康检查等用途
	App.HEAD("/ok", func(c *gin.Context) { c.String(http.StatusOK, "ok") })

	// 用户登录
	App.POST("/login", handler.Login)
}

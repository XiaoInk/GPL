/*
	Created by XiaoInk at 2021/04/23 16:23:17
	GitHub: https://github.com/XiaoInk
*/

package router

import (
	"net/http"

	"github.com/XiaoInk/GPL/handler"
	"github.com/XiaoInk/GPL/middleware"
	"github.com/gin-gonic/gin"
)

var App *gin.Engine

func init() {
	App = gin.Default()

	// 健康检查
	App.HEAD("/ok", func(c *gin.Context) { c.String(http.StatusOK, "ok") })

	// 登录
	App.POST("/login", handler.Login)

	// 需要授权
	Auth := App.Group("/", middleware.Authenticator)
	{
		Auth.POST("/logout", handler.Logout)
	}
}

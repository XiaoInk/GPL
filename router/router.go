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

	// 解决跨域问题
	App.Use(middleware.Cors)

	// 健康检查
	App.HEAD("/ok", func(c *gin.Context) { c.String(http.StatusOK, "ok") })

	// 登录
	App.POST("/login", handler.Login)

	// 需要授权
	Auth := App.Group("/", middleware.Authenticator)
	{
		// 注销
		Auth.POST("/logout", handler.Logout)

		// 添加域名
		Auth.POST("/domains", handler.AddDomain)
		// 获取域名列表
		Auth.GET("/domains", handler.GetDomains)
	}
}

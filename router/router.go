/*
	Created by XiaoInk at 2021/04/23 16:23:17
	GitHub: https://github.com/XiaoInk
*/

package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var App *gin.Engine

func init() {
	App = gin.Default()

	// 可作健康检查等用途
	App.Any("/ok", func(c *gin.Context) { c.String(http.StatusOK, "ok") })
}

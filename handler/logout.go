/*
	Created by XiaoInk at 2021/04/25 23:21:25
	GitHub: https://github.com/XiaoInk
*/

package handler

import (
	"net/http"

	"github.com/XiaoInk/GPL/logic"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	res := NewResponse()
	err := logic.Logout(c.GetHeader("Authorization"))
	if err != nil {
		res.SetMetaMsg("注销失败")
	} else {
		res.SetMeta(map[string]interface{}{
			"msg":    "注销成功",
			"status": http.StatusOK,
		})
	}

	c.JSON(http.StatusOK, res)
}

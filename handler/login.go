/*
	Created by XiaoInk at 2021/04/24 00:21:26
	GitHub: https://github.com/XiaoInk
*/

package handler

import (
	"net/http"

	"github.com/XiaoInk/GPL/logic"
	"github.com/gin-gonic/gin"
)

type LoginParams struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Bind(ctx *gin.Context, params interface{}) {
	res := NewResponse()
	ctx.AbortWithStatusJSON(http.StatusOK, res)
}

func Login(ctx *gin.Context) {
	var p LoginParams

	res := NewResponse()
	err := ctx.ShouldBind(&p)
	if err != nil {
		ctx.JSON(http.StatusOK, res)
		return
	}

	token, err := logic.Login(p.Username, p.Password)
	if err != nil {
		res.SetMetaMessage(err.Error())
		ctx.JSON(http.StatusOK, res)
		return
	}

	res.SetMeta(map[string]interface{}{
		"code":    http.StatusOK,
		"token":   token,
		"message": "登录成功",
	})
	ctx.JSON(http.StatusOK, res)
}

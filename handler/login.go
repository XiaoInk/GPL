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

type LoginReq struct {
	Username string `form:"username" json:"username" binding:"required,min=4,max=20"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=25"`
}

func Login(c *gin.Context) {
	var p LoginReq

	res := NewResponse()
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusOK, res)
		return
	}

	token, err := logic.Login(logic.LoginParams{Username: p.Username, Password: p.Password})
	if err != nil {
		res.SetMetaMsg(err.Error())
	} else {
		res.SetMeta(map[string]interface{}{
			"msg":    "登录成功",
			"status": http.StatusOK,
			"token":  token,
		})
	}

	c.JSON(http.StatusOK, res)
}

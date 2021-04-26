/*
	Created by XiaoInk at 2021/04/25 23:26:32
	GitHub: https://github.com/XiaoInk
*/

package middleware

import (
	"net/http"

	"github.com/XiaoInk/GPL/dao"
	"github.com/XiaoInk/GPL/handler"
	"github.com/gin-gonic/gin"
)

func Authenticator(c *gin.Context) {
	res := handler.NewResponse()
	userId, err := dao.GetTokenInt(c.GetHeader("Authorization"))
	if err != nil {
		res.SetMeta(map[string]interface{}{
			"msg":    "令牌无效",
			"status": http.StatusUnauthorized,
		})
		c.JSON(http.StatusOK, res)
		c.Abort()
		return
	}

	user := dao.NewUser()
	ok := user.IsExistById(userId)
	if !ok {
		res.SetMetaMsg("获取用户信息失败")
		c.JSON(http.StatusOK, res)
		c.Abort()
		return
	}

	c.Set("userId", user.ID)     // uint
	c.Set("userName", user.Name) // string
	c.Set("roleId", user.RoleID) // uint
}

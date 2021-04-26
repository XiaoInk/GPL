/*
	Created by XiaoInk at 2021/04/26 11:00:32
	GitHub: https://github.com/XiaoInk
*/

package handler

import (
	"net/http"

	"github.com/XiaoInk/GPL/logic"
	"github.com/gin-gonic/gin"
)

type AddDomainReq struct {
	Name string `form:"name" json:"name" binding:"required,fqdn"`
}

func AddDomain(c *gin.Context) {
	var p AddDomainReq

	res := NewResponse()
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusOK, res)
		return
	}

	if err := logic.AddDomain(p.Name, c.GetUint("userId"), c.GetString("userName")); err != nil {
		res.SetMetaMsg(err.Error())
	} else {
		res.SetMeta(map[string]interface{}{
			"msg":    "添加域名成功",
			"status": http.StatusCreated,
		})
	}

	c.JSON(http.StatusOK, res)
}

func GetDomains(c *gin.Context) {
	res := NewResponse()
	pager := NewPager()
	c.ShouldBindQuery(&pager)

	pager.Init()
	total, domains, err := logic.GetDomains(c.Query("name"), pager.Limit, pager.Offset, pager.Order)
	if err != nil {
		res.SetMetaMsg(err.Error())
	} else {
		res.SetData(map[string]interface{}{
			"total":       total,
			"domain_list": domains,
		})
		res.SetMeta(map[string]interface{}{
			"msg":    "获取域名列表成功",
			"status": http.StatusOK,
		})
	}

	c.JSON(http.StatusOK, res)
}

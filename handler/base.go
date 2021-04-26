/*
	Created by XiaoInk at 2021/04/24 09:18:18
	GitHub: https://github.com/XiaoInk
*/

package handler

import (
	"net/http"
)

type pager struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pagesize"`
	Order    string `form:"order"`
	Limit    int
	Offset   int
}

func NewPager() *pager {
	return &pager{
		Page:     1,
		PageSize: 10,
		Order:    "updated_at desc",
	}
}

func (p *pager) Init() {
	p.Limit = p.PageSize
	p.Offset = (p.Page - 1) * p.Limit
}

type response struct {
	Data interface{}            `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

func NewResponse() *response {
	return &response{
		Meta: map[string]interface{}{
			"msg":    "请求参数错误",
			"status": http.StatusBadRequest,
		},
	}
}

func (r *response) SetData(data interface{}) { r.Data = data }

func (r *response) SetMeta(meta map[string]interface{}) { r.Meta = meta }

func (r *response) SetMetaMsg(msg string) { r.Meta["msg"] = msg }

func (r *response) SetMetaStatus(code int) { r.Meta["status"] = code }

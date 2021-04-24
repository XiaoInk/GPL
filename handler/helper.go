/*
	Created by XiaoInk at 2021/04/24 09:18:18
	GitHub: https://github.com/XiaoInk
*/

package handler

import (
	"net/http"
)

type response struct {
	Data interface{}            `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

func NewResponse() *response {
	return &response{
		Meta: map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "请求参数错误",
		},
	}
}

func (r *response) SetData(data interface{}) { r.Data = data }

func (r *response) SetMeta(meta map[string]interface{}) { r.Meta = meta }

func (r *response) SetMetaCode(code int) { r.Meta["code"] = code }

func (r *response) SetMetaMessage(msg string) { r.Meta["message"] = msg }

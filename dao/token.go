/*
	Created by XiaoInk at 2021/04/23 23:55:14
	GitHub: https://github.com/XiaoInk
*/

package dao

import (
	"time"

	"github.com/XiaoInk/GPL/model"
)

type token struct{ Str string }

func NewToken(s string) *token {
	return &token{Str: model.Config.TokenPrefix + s}
}

func (t *token) SetCache(val interface{}, exp time.Duration) error {
	rdb, ctx := model.GetRdb()
	return rdb.Set(ctx, t.Str, val, exp*time.Second).Err()
}

func (t *token) GetCache() (int, error) {
	rdb, ctx := model.GetRdb()
	return rdb.Get(ctx, t.Str).Int()
}

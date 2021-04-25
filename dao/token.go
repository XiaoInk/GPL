/*
	Created by XiaoInk at 2021/04/23 23:55:14
	GitHub: https://github.com/XiaoInk
*/

package dao

import (
	"time"

	"github.com/XiaoInk/GPL/model"
	"github.com/XiaoInk/GPL/util"
)

func NewToken(s string) string { return util.MD5(s + util.RandString(4)) }

func SetToken(token string, val interface{}, exp time.Duration) error {
	rdb, ctx := model.GetRdb()
	return rdb.Set(ctx, model.Config.TokenPrefix+token, val, exp*time.Second).Err()
}

func GetTokenInt(token string) (int, error) {
	rdb, ctx := model.GetRdb()
	return rdb.Get(ctx, model.Config.TokenPrefix+token).Int()
}

func DelToken(token string) error {
	rdb, ctx := model.GetRdb()
	return rdb.Del(ctx, model.Config.TokenPrefix+token).Err()
}

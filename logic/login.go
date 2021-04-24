/*
	Created by XiaoInk at 2021/04/23 23:47:00
	GitHub: https://github.com/XiaoInk
*/

package logic

import (
	"errors"

	"github.com/XiaoInk/GPL/dao"
	"github.com/XiaoInk/GPL/util"
)

func Login(username, password string) (string, error) {
	user := dao.NewUser()
	user.Username = username
	user.Password = password

	ok := user.MustMatchUsernameAndPassword()
	if !ok {
		return "", errors.New("登录失败")
	}

	token := dao.NewToken(util.MD5(user.Username + util.RandString(4)))
	err := token.SetCache(user.ID, 60*60)
	if err != nil {
		return "", errors.New("缓存失败")
	}

	return token.Str, nil
}

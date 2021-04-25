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

type LoginParams struct{ Username, Password string }

func Login(p LoginParams) (string, error) {
	user := dao.NewUser()
	user.Username = p.Username
	user.Password = util.MD5(p.Password)

	ok := user.MatchUsernameAndPassword()
	if !ok {
		return "", errors.New("用户名或密码错误")
	}

	token := dao.NewToken(user.Username)
	err := dao.SetToken(token, user.ID, 60*60*8)
	if err != nil {
		return "", errors.New("缓存失败")
	}

	return token, nil
}

/*
	Created by XiaoInk at 2021/04/23 22:27:51
	GitHub: https://github.com/XiaoInk
*/

package dao

import (
	"github.com/XiaoInk/GPL/model"
	"github.com/XiaoInk/GPL/model/table"
)

type user table.User

func NewUser() *user { return &user{} }

// 验证用户名、密码是否匹配
func (u *user) MustMatchUsernameAndPassword() bool {
	return model.Getdb().Where("username = ? and password = ?", u.Username, u.Password).Take(&u).RowsAffected == 1
}

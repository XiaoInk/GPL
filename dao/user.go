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

func (u *user) GetUserById(userId int) bool {
	return model.Getdb().Where("id = ?", userId).Take(&u).RowsAffected == 1
}

// 匹配用户名、密码
func (u *user) MatchUsernameAndPassword() bool {
	return model.Getdb().Where("username = ? and password = ?", u.Username, u.Password).Take(&u).RowsAffected == 1
}

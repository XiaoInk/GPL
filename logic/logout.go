/*
	Created by XiaoInk at 2021/04/25 23:22:00
	GitHub: https://github.com/XiaoInk
*/

package logic

import "github.com/XiaoInk/GPL/dao"

func Logout(token string) error {
	return dao.DelToken(token)
}

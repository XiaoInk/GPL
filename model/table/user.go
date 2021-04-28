/*
	Created by XiaoInk at 2021/04/23 21:45:38
	GitHub: https://github.com/XiaoInk
*/

package table

import (
	"github.com/XiaoInk/GPL/util"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"size:128;not null;comment:姓名" json:"name"`
	Username  string `gorm:"size:128;not null;unique;comment:登录用户" json:"username"`
	Password  string `gorm:"size:128;not null;comment:登录密码" json:"-"`
	Email     string `gorm:"size:128;not null;unique;comment:邮件地址" json:"email"`
	Mobile    string `gorm:"size:11;comment:手机号码" json:"mobile"`
	Status    bool   `gorm:"default:1;comment:用户状态, 0 禁用、1 启用" json:"status"`
	RoleID    uint   `gorm:"default:7;comment:角色ID" json:"role_id"`
	RoleName  string `gorm:"size:128;default:普通用户;comment:角色名称" json:"role_name"`
	CreatedAt int    `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt int    `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at"`
}

func (u *User) Init(db *gorm.DB) {
	if db.Where("username = 'admin'").Take(&u).RowsAffected != 1 {
		u.Name = "Admin"
		u.Username = "admin"
		u.Password = util.MD5("admin@example.com")
		u.Email = "admin@example.com"
		u.RoleID = 1
		u.RoleName = "超级管理"

		db.Create(&u)
	}
}

/*
	Created by XiaoInk at 2021/04/26 10:42:05
	GitHub: https://github.com/XiaoInk
*/

package table

type Domain struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"size:128;not null;unique;comment:域名" json:"name"`
	Registrar string `gorm:"size:128;comment:注册商" json:"registrar"`
	RegTime   string `gorm:"comment:注册时间" json:"reg_time"`
	ExpTime   string `gorm:"comment:过期时间" json:"exp_time"`
	UserId    uint   `gorm:"not null;comment:创建者ID" json:"user_id"`
	UserName  string `gorm:"size:128;not null;comment:创建者名称" json:"user_name"`
	CreatedAt int    `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt int    `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at"`
}

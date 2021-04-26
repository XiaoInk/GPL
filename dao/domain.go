/*
	Created by XiaoInk at 2021/04/26 12:10:18
	GitHub: https://github.com/XiaoInk
*/

package dao

import (
	"github.com/XiaoInk/GPL/model"
	"github.com/XiaoInk/GPL/model/table"
)

type domain table.Domain

func NewDomain() *domain { return &domain{} }

func (d *domain) IsExistByName(name string) bool {
	return model.Getdb().Where("name = ?", name).Take(&d).RowsAffected == 1
}

func (d *domain) Create() error {
	return model.Getdb().Create(&d).Error
}

func (d *domain) GetDomainList(name string, limit, offset int, order string) (total int64, domains *[]domain, err error) {
	db := model.Getdb().Model(&domain{})
	switch name {
	case "":
		err = db.Count(&total).Limit(limit).Offset(offset).Order(order).Find(&domains).Error
	default:
		err = db.Where("name like ?", "%"+name+"%").Count(&total).Limit(limit).Offset(offset).Order(order).Find(&domains).Error
	}
	return
}

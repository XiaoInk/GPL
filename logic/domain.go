/*
	Created by XiaoInk at 2021/04/26 10:52:24
	GitHub: https://github.com/XiaoInk
*/

package logic

import (
	"errors"
	"time"

	"github.com/XiaoInk/GPL/dao"
	"github.com/XiaoInk/GPL/service"
)

// 添加域名
func AddDomain(name string, userId uint, userName string) error {
	domain := dao.NewDomain()
	ok := domain.IsExistByName(name)
	if ok {
		return errors.New("域名已存在")
	}

	info, err := service.GetDomainInfo(name)
	if err != nil {
		return errors.New("获取域名信息失败")
	}

	loc, _ := time.LoadLocation("Local")
	rt, _ := time.ParseInLocation("2006-01-02T15:04:05Z", info["registrationTime"], loc)
	et, _ := time.ParseInLocation("2006-01-02T15:04:05Z", info["expirationTime"], loc)

	domain.Name = name
	domain.Registrar = info["registrar"]
	domain.RegTime = rt.Format("2006-01-02 15:04:05")
	domain.ExpTime = et.Format("2006-01-02 15:04:05")
	domain.UserId = userId
	domain.UserName = userName

	if err := domain.Create(); err != nil {
		return errors.New("添加域名失败")
	}
	return nil
}

// 获取域名列表
func GetDomains(query string, limit, offset int, order string) (int64, interface{}, error) {
	domain := dao.NewDomain()
	total, domains, err := domain.GetDomainList(query, limit, offset, order)
	if err != nil {
		return 0, nil, errors.New("获取域名列表失败")
	}

	return total, domains, nil
}

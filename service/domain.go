/*
	Created by XiaoInk at 2021/04/26 10:44:35
	GitHub: https://github.com/XiaoInk
*/

package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"
)

func GetDomainInfo(name string) (map[string]string, error) {
	client := NewClient(5 * time.Second)
	resp, err := client.Get("https://api.devopsclub.cn/api/whoisquery?domain=" + name + "&type=json&standard=true")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var res struct {
		Code int
		Data map[string]map[string]string
	}
	_ = json.Unmarshal(body, &res)

	if res.Code != 0 {
		return nil, errors.New("请求域名信息失败")
	}

	return res.Data["data"], nil
}

/*
	Created by XiaoInk at 2021/04/23 20:47:26
	GitHub: https://github.com/XiaoInk
*/

package model

type config struct {
	TokenPrefix string
	MySQLUri    string
	TablePrefix string
	RedisUri    string
}

func NewConfig() *config {
	return &config{
		TokenPrefix: "GPL-",
		MySQLUri:    "root:123456@tcp(127.0.0.1:3306)/gpldb",
		TablePrefix: "gpl_",
		RedisUri:    "redis://:123456@127.0.0.1:6379/0",
	}
}

var Config *config

func init() { Config = NewConfig() }

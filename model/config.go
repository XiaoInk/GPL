/*
	Created by XiaoInk at 2021/04/23 18:24:30
	GitHub: https://github.com/XiaoInk
*/

package model

type config struct {
	MySQL string
	Redis string
}

func NewConfig() *config {
	return &config{
		MySQL: "root:123456@tcp(127.0.0.1:3306)/omsdb",
		Redis: "redis://:123456@127.0.0.1:6379/0",
	}
}

var Config *config

func init() { Config = NewConfig() }

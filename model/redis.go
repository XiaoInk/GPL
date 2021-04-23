/*
	Created by XiaoInk at 2021/04/23 19:18:11
	GitHub: https://github.com/XiaoInk
*/

package model

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func GetRdb() (*redis.Client, context.Context) {
	return rdb, context.Background()
}

func init() {
	opt, err := redis.ParseURL(Config.Redis)
	if err != nil {
		log.Fatalln("Redis.ParseURL.err", err.Error())
	}

	rdb = redis.NewClient(opt)
}

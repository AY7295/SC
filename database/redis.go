package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var RDB *redis.Client

func Redis(host, port, pwd string, code int) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pwd,
		DB:       code,
	})

	_, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		log.Println("redis数据库连接失败" + err.Error())
		return
	} else {
		fmt.Println("redis数据库连接成功")
	}

	RDB = rdb
	fmt.Println(RDB)
}

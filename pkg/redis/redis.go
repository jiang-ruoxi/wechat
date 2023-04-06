package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
	"wechat/config"
)

var RedisClient *redis.Client

func InitDB(conf config.Config) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr + ":" + conf.Redis.Port,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
		PoolSize: 100,
	})
	cxt, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := RedisClient.Ping(cxt).Result()
	if err != nil {
		log.Fatal("redis ping err::", err)
	}
	fmt.Println("redis init on port ", conf.Redis.Port)
}
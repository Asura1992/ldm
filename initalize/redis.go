package initalize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"ldm/common/config"
	"ldm/common/dao"
	"log"
)

//初始化redis
func InitRedis() {
	cfg := config.GlobalConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Address,
		DB:   cfg.Db,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("ping redis 服务失败:", err)
	}
	dao.Redis = client
}

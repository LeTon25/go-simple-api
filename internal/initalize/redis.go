package initalize

import (
	"context"
	"fmt"
	"go-simple-api/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.DB,       // use default DB
		PoolSize: r.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect to Redis", zap.Error(err))
	}

	global.Rdb = rdb

}

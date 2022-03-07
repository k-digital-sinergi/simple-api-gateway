package connection

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"simple-api-gateway/config"
	"sync"
	"time"
)

var once sync.Once
var redisInstance *redis.Client

func GetRedisConnection(ctx context.Context) *redis.Client {
	once.Do(func() {
		redisInstance = newRedisConnection(ctx)
	})

	return redisInstance
}

func newRedisConnection(ctx context.Context) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         config.Env.RedisAddress,
		Password:     config.Env.RedisPassword,
		PoolTimeout:  20 * time.Second,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	return rdb
}

func OpenRedisConnection() {
	GetRedisConnection(context.Background())
}

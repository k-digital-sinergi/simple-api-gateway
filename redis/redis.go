package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Redis interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
}

func Get(ctx context.Context, rds Redis, key string) (string, error) {
	data, err := rds.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			return "", err
		}
	}

	return data, nil
}

func Set(ctx context.Context, rds Redis, key, val string, exp time.Duration) error {
	_, err := rds.Set(ctx, key, val, exp).Result()
	if err != nil {
		return err
	}

	return nil
}

func Del(ctx context.Context, rds Redis, key string) error {
	_, err := rds.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	return nil
}

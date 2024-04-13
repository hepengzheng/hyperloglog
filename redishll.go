package gohll

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisHLL struct {
	rdb *redis.Client
}

var _ HyperLogLog = (*RedisHLL)(nil)

const redisKey = "hyperloglog-test"

func NewRedisHLL() *RedisHLL {
	return &RedisHLL{
		rdb: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}

func (rhll *RedisHLL) Add(ctx context.Context, s string) {
	rhll.rdb.PFAdd(ctx, redisKey, s)
}

func (rhll *RedisHLL) Count(ctx context.Context) int {
	return int(rhll.rdb.PFCount(ctx, redisKey).Val())
}

package db

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedis(ctx context.Context, addr string) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisClient{
		Client: client,
		Ctx:    ctx,
	}
}

func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.Client.Set(r.Ctx, key, value, expiration)
}

func (r *RedisClient) Get(key string) *redis.StringCmd {
	return r.Client.Get(r.Ctx, key)
}
func (r *RedisClient) HSet(key string, fields map[string]interface{}) *redis.IntCmd {
	return r.Client.HSet(r.Ctx, key, fields)
}
func (r *RedisClient) HGetAll(key string) *redis.MapStringStringCmd {
	return r.Client.HGetAll(r.Ctx, key)
}

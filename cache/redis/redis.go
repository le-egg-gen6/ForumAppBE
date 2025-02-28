package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
	ctx    context.Context
}

var Instance *RedisClient

func NewRedisClient(cfg *RedisConfig) *RedisClient {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.REDIS_HOST, cfg.REDIS_PORT),
		Password: cfg.REDIS_PASSWORD,
		DB:       cfg.REDIS_DATABASE,
	})

	ctx := context.Background()

	_, err := redisClient.Ping(ctx).Result()

	if err != nil {
		panic("Failed to connect to redis" + err.Error())
	}

	Instance = &RedisClient{
		Client: redisClient,
		ctx:    ctx,
	}
	return Instance
}

func GetRedisInstance() *RedisClient {
	return Instance
}

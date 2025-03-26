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

func InitializeRedis() {
	cfg, err := LoadRedisConfig()
	if err != nil {
		panic("Redis configuration not found")
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDatabase,
	})

	ctx := context.Background()

	_, err = redisClient.Ping(ctx).Result()

	if err != nil {
		panic("Failed to connect to redis" + err.Error())
	}

	Instance = &RedisClient{
		Client: redisClient,
		ctx:    ctx,
	}
}

func GetRedisInstance() *RedisClient {
	return Instance
}

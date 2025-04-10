package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
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

func SetWithTTL(key string, value interface{}, ttl time.Duration) error {
	return Instance.Client.Set(Instance.ctx, key, value, ttl).Err()
}

func SetWithoutTTL(key string, value interface{}) error {
	return Instance.Client.Set(Instance.ctx, key, value, 0).Err()
}

func Get[T any](key string) (T, error) {
	val, err := Instance.Client.Get(Instance.ctx, key).Result()
	if err != nil {
		var zero T
		return zero, err
	}
	var result T
	err = json.Unmarshal([]byte(val), &result)
	return result, err
}

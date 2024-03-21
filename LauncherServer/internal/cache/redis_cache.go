package cache

import (
	"LauncherServer/internal/utils"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisService struct {
	client *redis.Client
}

var RedisCacheService *RedisService

func InitCacheRedis(config utils.Config) {
	RedisCacheService = NewRedisService(
		redis.NewClient(&redis.Options{
			Addr:     config.RedisAddress,
			Password: config.RedisPassword,
			DB:       config.RedisDB,
		}))
}

func NewRedisService(client *redis.Client) *RedisService {
	return &RedisService{client: client}
}

func (r *RedisService) Set(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(context.Background(), key, value, expiration).Err()
}

func (r *RedisService) Get(key string) (string, error) {
	return r.client.Get(context.Background(), key).Result()
}

func (r *RedisService) Del(key string) error {
	return r.client.Del(context.Background(), key).Err()
}

func (r *RedisService) Exists(key string) (bool, error) {
	result, err := r.client.Exists(context.Background(), key).Result()
	return result > 0, err
}

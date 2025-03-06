package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/carlosblanco00/order-manager-api/internal/domain/model"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

type RedisRepository struct {
	Client *redis.Client
}

func (r *RedisRepository) SetKey(key string, value model.RedisModel, ttl time.Duration) error {

	valueBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Client.Set(Ctx, key, valueBytes, ttl).Err()
}

func (r *RedisRepository) CheckKey(key string) (bool, error) {
	val, err := r.Client.Exists(Ctx, key).Result()
	if err != nil {
		return false, err
	}
	return val > 0, nil
}

func (r *RedisRepository) GetKey(key string) (string, error) {
	return r.Client.Get(Ctx, key).Result()
}

func (r *RedisRepository) DeleteKey(key string) error {
	return r.Client.Del(Ctx, key).Err()
}

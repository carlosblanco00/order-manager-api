package model

import "time"

type RedisModel struct {
	IdempotencyKey string `json:"idempotency_key"`
	Status         string `json:"status"`
	Response       string `json:"response"`
}

type RedisGateway interface {
	SetKey(key string, value RedisModel, ttl time.Duration) error
	CheckKey(key string) (bool, error)
	GetKey(key string) (string, error)
	DeleteKey(key string) error
}

package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Client *redis.Client

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func InitRedisConnection() {
	Client = NewRedisClient()
	pong, err := Client.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("No se pudo conectar a Redis: %w", err)
	}
	log.Printf("Conexión exitosa: %v", pong)
}

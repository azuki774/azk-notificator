package repository

import (
	"azk-notificator/internal/model"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	Client *redis.Client
}

func (r *RedisClient) Push(q model.Queue) (err error) {
	return nil
}

func (r *RedisClient) Pop() (q model.Queue, err error) {
	return model.Queue{}, nil
}

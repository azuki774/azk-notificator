package repository

import (
	"azk-notificator/internal/model"

	"github.com/go-redis/redis"
)

type Redis struct {
	Client redis.Client
}

func (r *Redis) Push(q model.Queue) (err error) {
	return nil
}

func (r *Redis) Pop() (q model.Queue, err error) {
	return model.Queue{}, nil
}

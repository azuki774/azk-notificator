package repository

import (
	"azk-notificator/internal/model"
	"context"
	"encoding/json"

	"github.com/go-redis/redis"
)

const queueKey = "queue"

type Redis struct {
	Client   *redis.Client
	Capacity int
}

func (r *Redis) Push(ctx context.Context, q model.Queue) (err error) {
	pd, err := json.Marshal(q)
	if err != nil {
		return err
	}

	// Check Capacity
	cmd := r.Client.LLen(queueKey)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	num := int(cmd.Val())
	if num >= r.Capacity {
		return model.ErrOverCapacity
	}

	// push record
	err = r.Client.RPush(queueKey, string(pd)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Pop(ctx context.Context) (q model.Queue, err error) {
	return model.Queue{}, nil
}

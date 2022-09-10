package repository

import (
	"azk-notificator/internal/model"
	"context"
	"encoding/json"
	"errors"

	"github.com/go-redis/redis"
)

type queueKey struct{}

var rediskey queueKey

const defaultKey = "queue"

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
	cmd := r.Client.LLen(r.getKeyName(ctx))
	if cmd.Err() != nil {
		return cmd.Err()
	}
	num := int(cmd.Val())
	if num >= r.Capacity {
		return model.ErrOverCapacity
	}

	// push record
	err = r.Client.RPush(r.getKeyName(ctx), string(pd)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Pop(ctx context.Context) (q model.Queue, err error) {
	cmd := r.Client.LPop(r.getKeyName(ctx))
	if errors.Is(cmd.Err(), redis.Nil) {
		return model.Queue{}, model.ErrQueueNotFound
	}
	if err != nil {
		return model.Queue{}, err
	}

	allb, err := cmd.Bytes()
	if err != nil {
		return model.Queue{}, err
	}

	err = json.Unmarshal(allb, &q)
	if err != nil {
		return model.Queue{}, err
	}

	return q, nil
}

// getKeyName get redis-key name from ctx.
// If not exists from ctx, use default queueKey.
func (r *Redis) getKeyName(ctx context.Context) string {
	v := ctx.Value(rediskey)
	key, ok := v.(string)
	if !ok {
		return defaultKey
	}

	if key == "" {
		return defaultKey
	}
	return key
}

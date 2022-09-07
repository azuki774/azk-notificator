package mock

import (
	"azk-notificator/internal/model"
	"context"
)

type mockRedis struct {
	q   model.Queue
	err error
}

func NewMockRedisClient(err error) *mockRedis {
	return &mockRedis{err: err}
}

func (m *mockRedis) Push(ctx context.Context, q model.Queue) (err error) {
	return m.err
}

func (m *mockRedis) Pop(ctx context.Context) (q model.Queue, err error) {
	return m.q, m.err
}

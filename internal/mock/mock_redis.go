package mock

import "azk-notificator/internal/model"

type mockRedis struct {
	q   model.Queue
	err error
}

func NewMockRedisClient(err error) *mockRedis {
	return &mockRedis{err: err}
}

func (m *mockRedis) Push(q model.Queue) (err error) {
	return m.err
}

func (m *mockRedis) Pop() (q model.Queue, err error) {
	return m.q, m.err
}

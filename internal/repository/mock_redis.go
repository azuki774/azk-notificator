package repository

import "azk-notificator/internal/model"

type mockRedisClient struct {
	q   model.Queue
	err error
}

func NewMockRedisClient(err error) *mockRedisClient {
	return &mockRedisClient{err: err}
}

func (m *mockRedisClient) Push(q model.Queue) (err error) {
	return m.err
}

func (m *mockRedisClient) Pop() (q model.Queue, err error) {
	return m.q, m.err
}

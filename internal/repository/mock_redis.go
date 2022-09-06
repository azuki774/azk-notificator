package repository

import "azk-notificator/internal/model"

type mockRedisClient struct{}

func NewMockRedisClient() *mockRedisClient {
	return &mockRedisClient{}
}

func (m *mockRedisClient) Push(q model.Queue) (err error) {
	return err
}

func (m *mockRedisClient) Pop() (q model.Queue, err error) {
	return model.Queue{}, nil
}

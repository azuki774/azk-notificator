package mock

import (
	"azk-notificator/internal/model"
	"context"
)

type MockQueueClient struct {
	Err error
}
type MockSendClient struct {
	Err error
}

func (m *MockQueueClient) Push(ctx context.Context, q model.Queue) (err error) {
	return m.Err
}

func (m *MockQueueClient) Pop(ctx context.Context) (q model.Queue, err error) {
	return model.Queue{}, m.Err
}

func (m *MockSendClient) Send(ctx context.Context, q model.Queue) (err error) {
	return m.Err
}

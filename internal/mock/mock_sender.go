package mock

import "azk-notificator/internal/model"

type MockQueueClient struct {
	Err error
}
type MockSendClient struct {
	Err error
}

func (m *MockQueueClient) Push(q model.Queue) (err error) {
	return m.Err
}

func (m *MockQueueClient) Pop() (q model.Queue, err error) {
	return model.Queue{}, m.Err
}

func (m *MockSendClient) Send(q model.Queue) (err error) {
	return m.Err
}

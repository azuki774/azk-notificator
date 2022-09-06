package repository

import "azk-notificator/internal/model"

type mockSendGmailClient struct {
	errSend error
}

func NewmockSendGmailClient(err error) *mockSendGmailClient {
	return &mockSendGmailClient{errSend: err}
}

func (m *mockSendGmailClient) Send(q model.Queue) (err error) {
	return m.errSend
}

package mock

import (
	"azk-notificator/internal/model"
	"context"
)

type mockSendClientGmail struct {
	errSend error
}

func NewmockSendGmailClient(err error) *mockSendClientGmail {
	return &mockSendClientGmail{errSend: err}
}

func (m *mockSendClientGmail) Send(ctx context.Context, q model.Queue) (err error) {
	return m.errSend
}

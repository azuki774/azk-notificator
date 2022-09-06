package mock

import "azk-notificator/internal/model"

type mockSendClientGmail struct {
	errSend error
}

func NewmockSendGmailClient(err error) *mockSendClientGmail {
	return &mockSendClientGmail{errSend: err}
}

func (m *mockSendClientGmail) Send(q model.Queue) (err error) {
	return m.errSend
}

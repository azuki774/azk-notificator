package repository

type mockSendClientGmail struct {
	errSend error
}

func NewmockSendGmailClient(err error) *mockSendClientGmail {
	return &mockSendClientGmail{errSend: err}
}

func (m *mockSendClientGmail) Send(body string) (err error) {
	return m.errSend
}

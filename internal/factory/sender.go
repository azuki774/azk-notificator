package factory

import (
	"azk-notificator/internal/repository"
	"azk-notificator/internal/sender"
)

type SenderRunOption struct {
	QueueHost string
	QueuePort string
	QueuePass string
}

func NewSender(opts *SenderRunOption) *sender.Sender {
	l, _ := NewLogger()
	sender := sender.Sender{
		Logger:      l,
		QueueClient: NewRedis(opts.QueueHost, opts.QueuePort, opts.QueuePass),
		SendClient:  NewSendCli(),
	}
	return &sender
}

func NewSendCli() *sender.SendCli {
	return &sender.SendCli{SendClientGmail: repository.NewmockSendGmailClient(nil)} // TODO: mock
}

func NewSendClientGmail() repository.SendClientGmail {
	// TODO: Get from ENV
	return repository.SendClientGmail{}
}

package factory

import (
	"azk-notificator/internal/repository"
	"azk-notificator/internal/sender"

	"go.uber.org/zap"
)

type SenderRunOption struct {
	QueueHost string
	QueuePort string
	QueuePass string
}


func NewSender(opts *SenderRunOption) *sender.Sender {
	l, _ := NewLogger()
	sender := sender.Sender{
		Logger: l,
		queueClient: factory.NewRedisClient(opts.QueueHost, opts.QueuePort, opts.QueuePass)
		sendClient: factory.NewSendClient()
	}
	return &sender
}

func NewSendClient() *sender.SendClient {
	return &sender.SendClient{SendClientGmail: repository.NewmockSendGmailClient(nil)} // TODO: mock
}

func NewSendClientGmail() *repository.SendClientGmail {
	// TODO: Get from ENV
	return &repository.SendClientGmail{}
}

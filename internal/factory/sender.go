package factory

import (
	"azk-notificator/internal/repository"
	"azk-notificator/internal/sender"

	"go.uber.org/zap"
)

type SenderRunOption struct {
	ServerHost string
	ServerPort string
}

func NewSender(opts *SenderRunOption) *sender.Sender {
	l, _ := NewLogger()
	sender := sender.Sender{
		Logger:     l,
		SendClient: NewSendCli(l),
		ServerHost: opts.ServerHost,
		ServerPort: opts.ServerPort,
	}
	return &sender
}

func NewSendCli(l *zap.Logger) *sender.SendCli {
	return &sender.SendCli{
		SendClientOnlyLog: NewSendSencClientOnlyLog(l),
	}
}

func NewSendSencClientOnlyLog(l *zap.Logger) sender.SendClientOnlyLog {
	return sender.SendClientOnlyLog{Logger: l}
}

func NewSendClientGmail() repository.SendClientGmail {
	// TODO: Get from ENV
	return repository.SendClientGmail{}
}

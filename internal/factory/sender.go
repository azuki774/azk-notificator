package factory

import (
	"azk-notificator/internal/model"
	"azk-notificator/internal/repository"
	"azk-notificator/internal/sender"
)

func NewSendClient(einfo model.SendEmailInfo) *sender.SendClient {
	return &sender.SendClient{SendClientGmail: repository.NewmockSendGmailClient(nil)} // TODO: mock
}

func NewSendClientGmail(info model.SendEmailInfo) *repository.SendClientGmail {
	return &repository.SendClientGmail{}
}

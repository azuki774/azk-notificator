package factory

import (
	"azk-notificator/internal/model"
	"azk-notificator/internal/repository"
)

func NewSendGmailClient(info model.SendEmailInfo) *repository.SendGmailClient {
	return &repository.SendGmailClient{}
}

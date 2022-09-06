package sender

import "azk-notificator/internal/model"

type SendClient interface {
	Send(q model.Queue) (err error)
}

type SendClientEmail interface {
	Send(body string) (err error)
}

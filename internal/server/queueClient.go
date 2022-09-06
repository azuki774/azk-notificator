package server

import "azk-notificator/internal/model"

type QueueClient interface {
	Push(q model.Queue) (err error)
}

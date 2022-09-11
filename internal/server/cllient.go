package server

import (
	"azk-notificator/internal/model"
	"context"
)

type QueueClient interface {
	Push(ctx context.Context, q model.Queue) (err error)
	Pop(ctx context.Context) (q model.Queue, err error)
}

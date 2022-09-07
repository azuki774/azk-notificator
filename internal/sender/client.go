package sender

import (
	"azk-notificator/internal/model"
	"context"
)

type SendClient interface {
	Send(ctx context.Context, q model.Queue) (err error)
}

type SendCli struct {
	SendClientGmail SendClientEmail
}

type SendClientEmail interface {
	Send(ctx context.Context, q model.Queue) (err error)
}

type QueueClient interface {
	Push(ctx context.Context, q model.Queue) (err error)
	Pop(ctx context.Context) (q model.Queue, err error)
}

func (s *SendCli) Send(ctx context.Context, q model.Queue) (err error) {
	switch q.Kind {
	case model.QueueKindEmail:
		return s.SendClientGmail.Send(ctx, q)
	}
	return model.ErrQueueUnexpctedKind
}

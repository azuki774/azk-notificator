package sender

import (
	"azk-notificator/internal/model"
	"context"

	"go.uber.org/zap"
)

type SendClient interface {
	Send(ctx context.Context, q model.Queue) (err error)
}

type SendOneClient interface {
	Send(ctx context.Context, q model.Queue) (err error)
}

type QueueClient interface {
	Push(ctx context.Context, q model.Queue) (err error)
	Pop(ctx context.Context) (q model.Queue, err error)
}
type SendCli struct {
	SendClientOnlyLog SendOneClient
	SendClientGmail   SendOneClient
}

type SendClientOnlyLog struct {
	Logger *zap.Logger
}

func (s *SendCli) Send(ctx context.Context, q model.Queue) (err error) {
	switch q.Kind {
	case model.QueueKindOnlyLog:
		return s.SendClientOnlyLog.Send(ctx, q)
	case model.QueueKindEmail:
		return s.SendClientGmail.Send(ctx, q)
	}
	return model.ErrQueueUnexpctedKind
}

func (l SendClientOnlyLog) Send(ctx context.Context, q model.Queue) (err error) {
	l.Logger.Info("send only log", zap.String("from", q.From), zap.String("to", q.To), zap.Int("kind", int(q.Kind)), zap.String("title", q.Title), zap.String("body", q.Body))
	return nil
}

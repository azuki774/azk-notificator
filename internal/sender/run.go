package sender

import (
	"azk-notificator/internal/model"
	"context"
	"errors"

	"go.uber.org/zap"
)

type Sender struct {
	Logger      *zap.Logger
	queueClient QueueClient
	sendClient  SendClient
}

func (s *Sender) Start(ctx context.Context) (err error) {
	s.Logger.Info("sender start")

	q, err := s.queueClient.Pop()
	if err != nil {
		if !errors.Is(err, model.ErrQueueNotFound) {
			s.Logger.Error("failed to dequeue", zap.Error(err))
			return err
		}
	}

	err = s.sendClient.Send(q)
	if err != nil {
		s.Logger.Error("failed to send the notification", zap.Error(err))
	}

	return nil
}

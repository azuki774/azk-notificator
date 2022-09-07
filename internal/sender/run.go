package sender

import (
	"azk-notificator/internal/model"
	"errors"

	"go.uber.org/zap"
)

type Sender struct {
	Logger      *zap.Logger
	QueueClient QueueClient
	SendClient  SendClient
}

func (s *Sender) Run() (err error) {
	s.Logger.Info("sender start")

	q, err := s.QueueClient.Pop()
	if err != nil {
		if errors.Is(err, model.ErrQueueNotFound) {
			return nil
		}
		s.Logger.Error("failed to dequeue", zap.Error(err))
		return err
	}

	err = s.SendClient.Send(q)
	if err != nil {
		s.Logger.Error("failed to send the notification", zap.Error(err))
		return err
	}

	return nil
}

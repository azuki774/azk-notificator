package sender

import (
	"go.uber.org/zap"
)

type Sender struct {
	Logger      *zap.Logger
	queueClient QueueClient
}

func (s *Sender) Start() (err error) {
	s.Logger.Info("sender start")

	return nil
}

package sender

import "azk-notificator/internal/model"

type SendClient interface {
	Send(q model.Queue) (err error)
}

type SendCli struct {
	SendClientGmail SendClientEmail
}

type SendClientEmail interface {
	Send(q model.Queue) (err error)
}

type QueueClient interface {
	Push(q model.Queue) (err error)
	Pop() (q model.Queue, err error)
}

func (s *SendCli) Send(q model.Queue) (err error) {
	switch q.Kind {
	case model.QueueKindEmail:
		return s.SendClientGmail.Send(q)
	}
	return model.ErrQueueUnexpctedKind
}

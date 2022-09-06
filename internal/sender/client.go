package sender

import "azk-notificator/internal/model"

type SendClient interface {
	Send(q model.Queue) (err error)
}

type SendCli struct {
	SendClientGmail SendClientEmail
}

func (s *SendCli) Send(q model.Queue) (err error) {
	switch q.Kind {
	case model.QueueKindEmail:
		return s.SendClientGmail.Send(q.Body)
	}
	return model.ErrQueueUnexpctedKind
}

type SendClientEmail interface {
	Send(body string) (err error)
}

type QueueClient interface {
	Push(q model.Queue) (err error)
	Pop() (q model.Queue, err error)
}

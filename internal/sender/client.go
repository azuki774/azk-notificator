package sender

import "azk-notificator/internal/model"

type SendClient struct {
	SendClientGmail SendClientEmail
}

func (s *SendClient) Send(q model.Queue) (err error) {
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

package model

import (
	"encoding/json"
	"errors"
)

type Queue struct {
	Kind Queuekind
	Body string
}

type Queuekind int

var (
	ErrQueueUnexpctedKind = errors.New("unexpected queue kind")
	ErrQueueNotFound      = errors.New("queue not found")
)

const (
	QueueKindEmail    = Queuekind(1)
	QueueKindEmailStr = "email"
)

func (q *Queue) UnmarshalJSON(b []byte) error {
	type QueueJSON struct {
		Kind string `json:"kind"`
		Body string `json:"body"`
	}

	var tq QueueJSON
	err := json.Unmarshal(b, &tq)
	if err != nil {
		return err
	}

	switch tq.Kind {
	case QueueKindEmailStr:
		q.Kind = QueueKindEmail
	default:
		return ErrQueueUnexpctedKind
	}

	q.Body = tq.Body
	return nil
}

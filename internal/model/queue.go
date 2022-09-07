package model

import (
	"errors"
)

type Queue struct {
	From  string
	To    string
	Kind  Queuekind
	Title string
	Body  string
}

type Queuekind int

var (
	ErrQueueUnexpctedKind = errors.New("unexpected queue kind")
	ErrQueueNotFound      = errors.New("queue not found")
	ErrOverCapacity       = errors.New("over capacity")
)

const (
	QueueKindEmail    = Queuekind(1)
	QueueKindEmailStr = "email"
)

// func (q *Queue) UnmarshalJSON(b []byte) error {
// 	type QueueJSON struct {
// 		From  string `json:"from"`
// 		To    string `json:"to"`
// 		Kind  string `json:"kind"`
// 		Title string `json:"title"`
// 		Body  string `json:"body"`
// 	}

// 	var tq QueueJSON
// 	err := json.Unmarshal(b, &tq)
// 	if err != nil {
// 		return err
// 	}

// 	q.From = tq.From
// 	q.To = tq.To

// 	switch tq.Kind {
// 	case QueueKindEmailStr:
// 		q.Kind = QueueKindEmail
// 	default:
// 		return ErrQueueUnexpctedKind
// 	}

// 	q.Title = tq.Title
// 	q.Body = tq.Body
// 	return nil
// }

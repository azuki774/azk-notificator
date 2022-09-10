package model

import (
	"errors"
)

type Queue struct {
	From  string    `json:"from"`
	To    string    `json:"to"`
	Kind  Queuekind `json:"kind"`
	Title string    `json:"title"`
	Body  []byte    `json:"body"`
}

type EnqueueHeader struct {
	From  string `schema:"from"`
	To    string `schema:"to"`
	Title string `schema:"title"`
}

type Queuekind int

var (
	ErrQueueUnexpctedKind = errors.New("unexpected queue kind")
	ErrQueueNotFound      = errors.New("queue not found")
	ErrOverCapacity       = errors.New("over capacity")
)

const (
	QueueKindOnlyLog    = Queuekind(1)
	QueueKindOnlyLogStr = "logonly"
	QueueKindEmail      = Queuekind(2)
	QueueKindEmailStr   = "email"
)

// func (q *Queue) UnmarshalJSON(b []byte) error {
// 	type QueueTmp struct {
// 		From  string    `json:"from"`
// 		To    string    `json:"to"`
// 		Kind  Queuekind `json:"kind"`
// 		Title string    `json:"title"`
// 		Body  []byte    `json:"body"`
// 	}

// 	sb := string(b)
// 	repb := strings.ReplaceAll(sb, "\n", "<br>")
// 	fmt.Println(repb)

// 	var qtmp QueueTmp
// 	err := json.Unmarshal([]byte(repb), &qtmp)
// 	if err != nil {
// 		return err
// 	}

// 	q.From = qtmp.From
// 	q.To = qtmp.To
// 	q.Kind = qtmp.Kind
// 	q.Title = qtmp.Title
// 	q.Body = strings.ReplaceAll(qtmp.Body, "<br>", "\n")

// 	return nil
// }

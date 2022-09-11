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

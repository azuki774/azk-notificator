package model

type Queue struct {
	kind Queuekind
	body string
}

type Queuekind int

const (
	QueueKindEmail = int(1)
)

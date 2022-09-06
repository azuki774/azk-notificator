package factory

import (
	"azk-notificator/internal/repository"
	"azk-notificator/internal/server"
)

type ServerRunOption struct {
	Host      string
	Port      string
	QueueHost string
	QueuePort string
	QueuePass string
}

func NewServer(opts *ServerRunOption) (*server.Server, error) {
	l, err := NewLogger()
	if err != nil {
		return nil, err
	}
	if opts.Port == "" {
		opts.Port = "80" // default value
	}

	// rdb := NewRedisClient(opts.QueueHost, opts.QueuePort, opts.QueuePort)
	rdb := repository.NewMockRedisClient()
	return &server.Server{Logger: l, Host: opts.Host, Port: opts.Port, QueueClient: rdb}, nil
}

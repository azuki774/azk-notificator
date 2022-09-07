package factory

import (
	"azk-notificator/internal/repository"
	"net"

	"github.com/go-redis/redis"
)

func NewRedis(host string, port string, password string) *repository.Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(host, port),
		Password: password,
		DB:       0, // use default DB
	})

	return &repository.Redis{Client: rdb}
}

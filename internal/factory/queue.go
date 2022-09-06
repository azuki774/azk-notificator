package factory

import (
	"net"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(host string, port string, password string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(host, port),
		Password: password,
		DB:       0, // use default DB
	})

	return rdb
}

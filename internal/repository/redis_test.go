package repository

import (
	"azk-notificator/internal/model"
	"context"
	"net"
	"testing"

	"github.com/go-redis/redis"
)

func TestRedis_Push(t *testing.T) {
	type fields struct {
		Client   *redis.Client
		Capacity int
	}
	type args struct {
		ctx context.Context
		q   model.Queue
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Client: redis.NewClient(&redis.Options{
					Addr:     net.JoinHostPort("localhost", "6379"),
					Password: "",
					DB:       0, // use default DB
				}),
				Capacity: 10,
			},
			args:    args{q: model.Queue{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				Client:   tt.fields.Client,
				Capacity: tt.fields.Capacity,
			}
			if err := r.Push(tt.args.ctx, tt.args.q); (err != nil) != tt.wantErr {
				t.Errorf("Redis.Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

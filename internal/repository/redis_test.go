package repository

import (
	"azk-notificator/internal/model"
	"context"
	"net"
	"reflect"
	"testing"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

var testBody string = "test\nてすと\n試験"

func TestMain(m *testing.M) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort("localhost", "6379"),
		Password: "",
		DB:       0, // use default DB
	})

	poptest1str := `{"to": "test123+abcde@gmail.com","kind": 1,"title": "title","body":"test\nてすと\n試験"}`
	err := redisClient.RPush("poptest1", poptest1str).Err()
	if err != nil {
		panic(err)
	}
	m.Run()
}

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
				Client:   redisClient,
				Capacity: 10,
			},
			args: args{
				ctx: context.WithValue(context.Background(), rediskey, "test1"),
				q: model.Queue{
					To:    "abcde@xyz.com",
					Kind:  model.QueueKindEmail,
					Title: "title",
					Body:  testBody,
				},
			},
			wantErr: false,
		},
		{
			name: "capacity over",
			fields: fields{
				Client:   redisClient,
				Capacity: 1,
			},
			args: args{
				ctx: context.WithValue(context.Background(), rediskey, "poptest1"),
				q: model.Queue{
					To:    "abcde@xyz.com",
					Kind:  model.QueueKindEmail,
					Title: "title",
					Body:  testBody,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				Client:   tt.fields.Client,
				Capacity: tt.fields.Capacity,
			}
			if err := r.Push(tt.args.ctx, tt.args.q); (err != nil) != tt.wantErr {
				t.Errorf("Redis.Push() error = %+v, wantErr %+v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_getKeyName(t *testing.T) {
	type fields struct {
		Client   *redis.Client
		Capacity int
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "default",
			fields: fields{
				Client:   redisClient,
				Capacity: 10,
			},
			args: args{
				ctx: context.Background(),
			},
			want: "queue",
		},
		{
			name: "select name",
			fields: fields{
				Client:   redisClient,
				Capacity: 10,
			},
			args: args{
				ctx: context.WithValue(context.Background(), rediskey, "getkeyname"),
			},
			want: "getkeyname",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				Client:   tt.fields.Client,
				Capacity: tt.fields.Capacity,
			}
			if got := r.getKeyName(tt.args.ctx); got != tt.want {
				t.Errorf("Redis.getKeyName() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestRedis_Pop(t *testing.T) {
	type fields struct {
		Client   *redis.Client
		Capacity int
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantQ   model.Queue
		wantErr bool
	}{
		{
			name: "found",
			fields: fields{
				Client:   redisClient,
				Capacity: 10,
			},
			args: args{
				ctx: context.WithValue(context.Background(), rediskey, "poptest1"),
			},
			wantQ: model.Queue{
				To:    "test123+abcde@gmail.com",
				Kind:  1,
				Title: "title",
				Body:  testBody,
			},
			wantErr: false,
		},
		{
			name: "not found",
			fields: fields{
				Client:   redisClient,
				Capacity: 10,
			},
			args: args{
				ctx: context.WithValue(context.Background(), rediskey, "poptestXXX"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				Client:   tt.fields.Client,
				Capacity: tt.fields.Capacity,
			}
			gotQ, err := r.Pop(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Redis.Pop() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotQ, tt.wantQ) {
				t.Errorf("Redis.Pop() = %+v, want %+v", gotQ, tt.wantQ)
			}
		})
	}
}

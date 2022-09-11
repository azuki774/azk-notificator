package server

import (
	"azk-notificator/internal/mock"
	"azk-notificator/internal/model"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"go.uber.org/zap"
)

var l *zap.Logger

var (
	testBody []byte = []byte{116, 101, 115, 116, 10, 227, 129, 166, 227, 129, 153, 227, 129, 168, 10, 232, 169, 166, 233, 168, 147}

	testdata1Str string = `{"to": "test123+abcde@gmail.com",
	"kind": 1,"title": "title",
	"body":[116, 101, 115, 116, 10, 227, 129, 166, 227, 129, 153, 227, 129, 168, 10, 232, 169, 166, 233, 168, 147]}`
	testdata1ReaderCloser io.ReadCloser
)

func TestMain(m *testing.M) {
	testdata1Reader := strings.NewReader(testdata1Str)
	testdata1ReaderCloser = ioutil.NopCloser(testdata1Reader)

	l, _ = zap.NewProduction()
	m.Run()
}

func TestServer_Enqueue(t *testing.T) {
	type fields struct {
		Logger      *zap.Logger
		QueueClient QueueClient
		Host        string
		Port        string
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
			name: "enqueue",
			fields: fields{
				Logger:      l,
				QueueClient: mock.NewMockRedisClient(nil),
			},
			args: args{
				ctx: context.Background(),
				q: model.Queue{
					To:   "testtest@abc.com",
					Kind: 1,
					Body: testBody,
				},
			},
			wantErr: false,
		},
		{
			name: "failed",
			fields: fields{
				Logger:      l,
				QueueClient: mock.NewMockRedisClient(errors.New("error")),
			},
			args: args{
				ctx: context.Background(),
				q: model.Queue{
					To:   "testtest@abc.com",
					Kind: 1,
					Body: testBody,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Logger:      tt.fields.Logger,
				QueueClient: tt.fields.QueueClient,
				Host:        tt.fields.Host,
				Port:        tt.fields.Port,
			}
			if err := s.Enqueue(tt.args.ctx, tt.args.q); (err != nil) != tt.wantErr {
				t.Errorf("Server.Enqueue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

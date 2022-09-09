package server

import (
	"azk-notificator/internal/mock"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"go.uber.org/zap"
)

var l *zap.Logger

var (
	testdata1Str          string = `{"to": "test123+abcde@gmail.com","kind": 1,"title": "title","body":"test\nてすと\n試験"}`
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
		r   *http.Request
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
				r: &http.Request{
					Body: testdata1ReaderCloser,
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
				r: &http.Request{
					Body: testdata1ReaderCloser,
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
			if err := s.Enqueue(tt.args.ctx, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Server.Enqueue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

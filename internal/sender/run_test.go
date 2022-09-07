package sender

import (
	"azk-notificator/internal/mock"
	"errors"
	"testing"

	"go.uber.org/zap"
)

var l *zap.Logger

func TestMain(m *testing.M) {
	l, _ = zap.NewProduction()
	m.Run()
}
func TestSender_Run(t *testing.T) {
	type fields struct {
		Logger      *zap.Logger
		QueueClient QueueClient
		SendClient  SendClient
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Logger:      l,
				QueueClient: &mock.MockQueueClient{},
				SendClient:  &mock.MockSendClient{},
			},
			wantErr: false,
		},
		{
			name: "failed dequeue",
			fields: fields{
				Logger:      l,
				QueueClient: &mock.MockQueueClient{Err: errors.New("error")},
				SendClient:  &mock.MockSendClient{},
			},
			wantErr: true,
		},
		{
			name: "failed send notification",
			fields: fields{
				Logger:      l,
				QueueClient: &mock.MockQueueClient{},
				SendClient:  &mock.MockSendClient{Err: errors.New("error")},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sender{
				Logger:      tt.fields.Logger,
				QueueClient: tt.fields.QueueClient,
				SendClient:  tt.fields.SendClient,
			}
			if err := s.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Sender.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

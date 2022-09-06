package sender

import (
	"azk-notificator/internal/mock"
	"azk-notificator/internal/model"
	"errors"
	"testing"
)

func TestSendCli_Send(t *testing.T) {
	type fields struct {
		SendClientGmail SendClientEmail
	}
	type args struct {
		q model.Queue
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "send email",
			fields:  fields{SendClientGmail: mock.NewmockSendGmailClient(nil)},
			args:    args{q: model.Queue{Kind: model.QueueKindEmail}},
			wantErr: false,
		},
		{
			name:    "send email failed",
			fields:  fields{SendClientGmail: mock.NewmockSendGmailClient(errors.New("error"))},
			args:    args{q: model.Queue{Kind: model.QueueKindEmail}},
			wantErr: true,
		},
		{
			name:    "unknown queue kind (send email)",
			fields:  fields{SendClientGmail: mock.NewmockSendGmailClient(nil)},
			args:    args{q: model.Queue{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SendCli{
				SendClientGmail: tt.fields.SendClientGmail,
			}
			if err := s.Send(tt.args.q); (err != nil) != tt.wantErr {
				t.Errorf("SendCli.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

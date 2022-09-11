package telemetry

import (
	"context"
	"testing"
)

func TestGetSpanIDWithCtx(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{ctx: NewCtxWithSpanID()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSpanIDWithCtx(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSpanIDWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (!tt.wantErr) && (got == "") {
				t.Errorf("GetSpanIDWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}

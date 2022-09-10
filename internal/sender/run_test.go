package sender

import (
	"testing"

	"go.uber.org/zap"
)

var l *zap.Logger

func TestMain(m *testing.M) {
	l, _ = zap.NewProduction()
	m.Run()
}

package telemetry

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type spanID string

const spanIDStr spanID = "spanID"

func GenUUID() uuid.UUID {
	uuidObj, _ := uuid.NewRandom()
	return uuidObj
}

func NewCtxWithSpanID(id uuid.UUID) context.Context {
	ctx := context.WithValue(context.Background(), spanIDStr, id.String)
	return ctx
}

func GetSpanIDWithCtx(ctx context.Context) (string, error) {
	v := ctx.Value(spanIDStr)

	spanID, ok := v.(string)
	if !ok {
		return "", errors.New("failed to get span-id")
	}

	return spanID, nil
}

func LoggerWithSpanID(ctx context.Context, l *zap.Logger) *zap.Logger {
	v, err := GetSpanIDWithCtx(ctx)
	if err != nil {
		l.Warn("failed to get span-id")
		return l
	}
	return l.With(zap.String("span-id", v))
}

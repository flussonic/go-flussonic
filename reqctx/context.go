package reqctx

import (
	"context"

	"github.com/google/uuid"
)

type TraceIDContextType string

const TraceIDContextKey = TraceIDContextType("trace_id")

func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDContextKey, traceID)
}

func WithNewTraceID(ctx context.Context) context.Context {
	return WithTraceID(ctx, uuid.NewString())
}

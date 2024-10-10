package logging

import (
	"context"

	"go.uber.org/zap"
)

// Context key is a key used to store the logger in the context.
type ContextKey struct{}

// GetContextKey returns the context key.
func GetContextKey() interface{} {
	return ContextKey{}
}

// WithLogger adds a logger to the context.
func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, GetContextKey(), logger)
}

// FromContext returns the logger from the context.
func FromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(GetContextKey()).(*zap.Logger); ok {
		return logger
	}

	// Return a new logger if none is found.
	return NewLogger()
}

package logging_test

import (
	"context"
	"testing"

	"github.com/nicklasfrahm-dev/appkit/logging"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestWithLogger(t *testing.T) {
	// Arrange.
	t.Parallel()

	logger := zap.NewNop()

	// Act.
	ctx := logging.WithLogger(context.TODO(), logger)
	retrievedLogger, ok := ctx.Value(logging.GetContextKey()).(*zap.Logger)

	// Assert.
	assert.True(t, ok, "should add the logger to the context")
	assert.Equal(t, logger, retrievedLogger, "should add the logger to the context")
}

func TestGetLoggerWithLogger(t *testing.T) {
	// Arrange.
	t.Parallel()

	logger := zap.NewNop()

	// Act.
	ctx := context.WithValue(context.TODO(), logging.GetContextKey(), logger)
	retrievedLogger := logging.FromContext(ctx)

	// Assert.
	assert.Equal(t, logger, retrievedLogger, "should retrieve the logger from the context")
}

func TestGetLoggerWithDefault(t *testing.T) {
	// Arrange.
	t.Parallel()

	// Act.
	logger := logging.FromContext(context.TODO())

	// Assert.
	assert.NotNil(t, logger, "should retrieve a default logger from the context")
}

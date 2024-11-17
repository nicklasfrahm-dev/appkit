package logging_test

import (
	"testing"

	"github.com/nicklasfrahm-dev/appkit/logging"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestWithDefaultLogLevel(t *testing.T) {
	// Arrange.
	t.Parallel()

	// Act.
	logger := logging.NewLogger()

	// Assert.
	assert.Equal(t, zap.InfoLevel, logger.Level(), "should default to the info log level")
}

func TestWithCustomLogLevel(t *testing.T) {
	// Arrange.
	t.Setenv("LOG_LEVEL", "debug")

	// Act.
	logger := logging.NewLogger()

	// Assert.
	assert.Equal(t, zap.DebugLevel, logger.Level(), "should set the log level to debug")
}

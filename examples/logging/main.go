package main

import (
	"fmt"

	"github.com/nicklasfrahm-dev/appkit/logging"
	"go.uber.org/zap"
)

func main() {
	logger := logging.NewLogger()

	port := 8080

	// Don't do this.
	logger.Info(fmt.Sprintf("Starting HTTP server on port %d", port))
	logger.Sugar().Infof("Starting HTTP server on port %d", port)

	// Do this instead.
	logger.Info("Starting HTTP server", zap.Int("port", port))

	printLog(logger)
}

// This is how you can pass the logger around.
func printLog(logger *zap.Logger) {
	logger.Info("This is a log message", zap.String("key", "value"))
}

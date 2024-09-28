# App Kit

This repository provides a set of useful libraries that can be used to quickly build a Go applications.

```shell
go get github.com/nicklasfrahm-dev/appkit
```

## Logging

This package provides a simple logging library based on [zap][github-zap], which may be configured at runtime using the environment variables shown below. The logger is optimized for structured logging using [Loki][github-loki] as a log aggregation system. By default the logger will use a production-ready configuration, but it can be configured to output logs in a more human-readable format for development purposes.

| Environment variable | Description                      | Default | Allowed values                            |
| -------------------- | -------------------------------- | ------- | ----------------------------------------- |
| `LOG_LEVEL`          | The minimum log level to output. | `info`  | `debug`, `info`, `warn`, `error`, `fatal` |
| `LOG_FORMAT`         | The log format to use.           | `json`  | `json`, `console`                         |

### Example

```go
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
```

## License

This project is and will always be licensed under the MIT License. See the [LICENSE.md](LICENSE.md) file for more information.

[github-zap]: https://github.com/uber-go/zap
[github-loki]: https://github.com/grafana/loki

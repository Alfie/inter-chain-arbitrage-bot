# Developer Documentation

## Logging
This project uses the [logrus](https://pkg.go.dev/github.com/sirupsen/logrus@v1.8.1 "Logrus Documentation") package to log events to the CLI. Set the desired log level using the global *LogLevel* variable. In the *development* environment, the logger reports events with function traceback.

# Developer Documentation

## Project Structure
The organization of the repository is based on the unofficial Golang [project layout](https://github.com/golang-standards/project-layout "Project Layout Standard") standard.

## Logging
This project uses the [logrus](https://pkg.go.dev/github.com/sirupsen/logrus@v1.8.1 "Logrus Documentation") package to log events to the CLI. Set the desired log level using the global *LogLevel* variable. In the *development* environment, the logger reports events with function traceback.

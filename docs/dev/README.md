# Developer Documentation

## Project Structure
The organization of this repository is based on the unofficial Golang [project-layout](https://github.com/golang-standards/project-layout "golang-standards project-layout") standard.

## Command Line Interface
The CLI is build using the [cobra](https://github.com/spf13/cobra "spf13 cobra") package.  According to the recommendation, the root file and all commands are placed in the *cmd/*  directory. The user input, configuration file and other environment variables are managed using the [viper](https://github.com/spf13/viper "spf13 viper") package.

## Logging
The project uses the [logrus](https://github.com/sirupsen/logrus "sirupsen logrus") package to log events to the CLI. Use the configuration file to set the *log level*, activate *function tracing* or specify the *format*.

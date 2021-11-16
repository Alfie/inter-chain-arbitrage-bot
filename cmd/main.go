package main

import (
	log "github.com/sirupsen/logrus"
	config "github.com/Sen-Com/inter-chain-arbitrage-bot/pkg/config"
)

func main() {
	// Mock initialize config variables
	// Should be done by the parser
	config.Environment = "development"
	config.LogLevel = "trace"

	init_logger()
}

func init_logger() {
	// Set logging format to json
	log.SetFormatter(&log.JSONFormatter{})

	// Activate/ Deactivate function specific reporting
	switch config.Environment {
	case "development":
		log.SetReportCaller(true) // Adds significant overhead
	case "production":
		log.SetReportCaller(false)
	}

	// Set the logging level
	switch config.LogLevel {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	}
}

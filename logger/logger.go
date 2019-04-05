package logger

import "github.com/sirupsen/logrus"

// Log -
var Log = new()

func new() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger
}

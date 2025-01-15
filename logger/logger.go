package logger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func New() *logrus.Logger {
	logFilePath := "debug.log" // TODO: change this depending on the env and maybe also config

	logger := logrus.New()
	rotatingLogger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}

	logger.SetOutput(rotatingLogger)

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	return logger
}

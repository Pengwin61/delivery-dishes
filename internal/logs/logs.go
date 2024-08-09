package logs

import (
	"go.uber.org/zap"
)

type Log struct {
	Logger *zap.Logger
}

func InitLogs() *Log {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return &Log{Logger: logger}
}

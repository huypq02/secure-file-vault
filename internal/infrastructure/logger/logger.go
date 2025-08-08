package logger

import "go.uber.org/zap"

var Log *zap.Logger

func Init() {
	Log, _ = zap.NewProduction()
}

package logs

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	rawLogger, _ := zap.NewDevelopment() // ใช้ zap.NewProduction() สำหรับ production
	// rawLogger = rawLogger.WithOptions(zap.AddCallerSkip(1))
	logger = rawLogger.Sugar()
}

func Info(msg string, fields ...interface{}) {
	logger.Infow(msg, fields...)
}

func Debug(msg string, fields ...interface{}) {
	logger.Debugw(msg, fields...)
}

func Error(msg string, fields ...interface{}) {
	logger.Errorw(msg, fields...)
}

func Fatal(msg string, fields ...interface{}) {
	logger.Fatalw(msg, fields...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}
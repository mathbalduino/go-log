package loggerCLI

import (
	"fmt"
)

func (l *loggerCLI) Trace(format string, args ...interface{}) LoggerCLI {
	logger := nestLogger(l)
	logger.baseLogger.Trace(fmt.Sprintf(format, args...))
	return logger
}
func (l *loggerCLI) Debug(format string, args ...interface{}) LoggerCLI {
	logger := nestLogger(l)
	logger.baseLogger.Debug(fmt.Sprintf(format, args...))
	return logger
}
func (l *loggerCLI) Info(format string, args ...interface{}) LoggerCLI {
	logger := nestLogger(l)
	logger.baseLogger.Info(fmt.Sprintf(format, args...))
	return logger
}
func (l *loggerCLI) Warn(format string, args ...interface{}) LoggerCLI {
	logger := nestLogger(l)
	logger.baseLogger.Warn(fmt.Sprintf(format, args...))
	return logger
}
func (l *loggerCLI) Error(format string, args ...interface{}) LoggerCLI {
	logger := nestLogger(l)
	logger.baseLogger.Error(fmt.Sprintf(format, args...))
	return logger
}
func (l *loggerCLI) Fatal(format string, args ...interface{}) {
	logger := nestLogger(l)
	logger.baseLogger.Fatal(fmt.Sprintf(format, args...))
}
func (l *loggerCLI) ErrorFrom(e error) LoggerCLI {
	logger := nestLogger(l)
	logger.baseLogger.ErrorFrom(e)
	return logger
}
func (l *loggerCLI) FatalFrom(e error) {
	logger := nestLogger(l)
	logger.baseLogger.FatalFrom(e)
}

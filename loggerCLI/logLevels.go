package loggerCLI

import (
	"fmt"
)

func (l *LoggerCLI) Trace(format string, args ...interface{}) *LoggerCLI {
	baseLogger := nestLogger(l)
	baseLogger.Trace(fmt.Sprintf(format, args...))
	return (*LoggerCLI)(baseLogger)
}
func (l *LoggerCLI) Debug(format string, args ...interface{}) *LoggerCLI {
	baseLogger := nestLogger(l)
	baseLogger.Debug(fmt.Sprintf(format, args...))
	return (*LoggerCLI)(baseLogger)
}
func (l *LoggerCLI) Info(format string, args ...interface{}) *LoggerCLI {
	baseLogger := nestLogger(l)
	baseLogger.Info(fmt.Sprintf(format, args...))
	return (*LoggerCLI)(baseLogger)
}
func (l *LoggerCLI) Warn(format string, args ...interface{}) *LoggerCLI {
	baseLogger := nestLogger(l)
	baseLogger.Warn(fmt.Sprintf(format, args...))
	return (*LoggerCLI)(baseLogger)
}
func (l *LoggerCLI) Error(format string, args ...interface{}) *LoggerCLI {
	baseLogger := nestLogger(l)
	baseLogger.Error(fmt.Sprintf(format, args...))
	return (*LoggerCLI)(baseLogger)
}
func (l *LoggerCLI) Fatal(format string, args ...interface{}) {
	baseLogger := nestLogger(l)
	baseLogger.Fatal(fmt.Sprintf(format, args...))
}
func (l *LoggerCLI) ErrorFrom(e error) *LoggerCLI {
	baseLogger := nestLogger(l)
	baseLogger.ErrorFrom(e)
	return (*LoggerCLI)(baseLogger)
}
func (l *LoggerCLI) FatalFrom(e error) {
	baseLogger := nestLogger(l)
	baseLogger.FatalFrom(e)
}

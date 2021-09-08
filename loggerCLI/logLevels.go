package loggerCLI

import (
	logger "gitlab.com/loxe-tools/go-log"
)

// TODO: criar standalone log functions, usando o logger default

func (l *LoggerCLI) Trace(msg string, adHocFields ...logger.LogFields) *LoggerCLI {
	baseLogger := nestLogger(l)
	baseLogger.Trace(msg, adHocFields...)
	return (*LoggerCLI)(baseLogger)
}
func (l *LoggerCLI) Debug(msg string, adHocFields ...logger.LogFields) *LoggerCLI {
	baseLogger := nestLogger(l)
	baseLogger.Debug(msg, adHocFields...)
	return (*LoggerCLI)(baseLogger)
}
func (l *LoggerCLI) Info(msg string, adHocFields ...logger.LogFields) *LoggerCLI {
	baseLogger := nestLogger(l)
	baseLogger.Info(msg, adHocFields...)
	return (*LoggerCLI)(baseLogger)
}
func (l *LoggerCLI) Warn(msg string, adHocFields ...logger.LogFields) *LoggerCLI {
	baseLogger := nestLogger(l)
	baseLogger.Warn(msg, adHocFields...)
	return (*LoggerCLI)(baseLogger)
}
func (l *LoggerCLI) Error(msg string, adHocFields ...logger.LogFields) *LoggerCLI {
	baseLogger := nestLogger(l)
	baseLogger.Error(msg, adHocFields...)
	return (*LoggerCLI)(baseLogger)
}
func (l *LoggerCLI) Fatal(msg string, adHocFields ...logger.LogFields) *LoggerCLI {
	baseLogger := nestLogger(l)
	baseLogger.Fatal(msg, adHocFields...)
	return (*LoggerCLI)(baseLogger)
}
func (l *LoggerCLI) ErrorFrom(e error, adHocFields ...logger.LogFields) *LoggerCLI {
	baseLogger := nestLogger(l)
	baseLogger.ErrorFrom(e, adHocFields...)
	return (*LoggerCLI)(baseLogger)
}
func (l *LoggerCLI) FatalFrom(e error, adHocFields ...logger.LogFields) *LoggerCLI {
	baseLogger := nestLogger(l)
	baseLogger.FatalFrom(e, adHocFields...)
	return (*LoggerCLI)(baseLogger)
}

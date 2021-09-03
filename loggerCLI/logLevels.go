package loggerCLI

import logger "gitlab.com/loxe-tools/go-log"

func (l *LoggerCLI) Trace(msg string, adHocFields ...logger.LogFields) *LoggerCLI {
	baseLogger := (*logger.Logger)(l)
	baseLogger.Trace(msg, adHocFields...)
	return nestLogger(baseLogger)
}
func (l *LoggerCLI) Debug(msg string, adHocFields ...logger.LogFields) *LoggerCLI {
	baseLogger := (*logger.Logger)(l)
	baseLogger.Debug(msg, adHocFields...)
	return nestLogger(baseLogger)
}
func (l *LoggerCLI) Info(msg string, adHocFields ...logger.LogFields) *LoggerCLI {
	baseLogger := (*logger.Logger)(l)
	baseLogger.Info(msg, adHocFields...)
	return nestLogger(baseLogger)
}
func (l *LoggerCLI) Warn(msg string, adHocFields ...logger.LogFields) *LoggerCLI {
	baseLogger := (*logger.Logger)(l)
	baseLogger.Warn(msg, adHocFields...)
	return nestLogger(baseLogger)
}
func (l *LoggerCLI) Error(msg string, adHocFields ...logger.LogFields) *LoggerCLI {
	baseLogger := (*logger.Logger)(l)
	baseLogger.Error(msg, adHocFields...)
	return nestLogger(baseLogger)
}
func (l *LoggerCLI) Fatal(msg string, adHocFields ...logger.LogFields) *LoggerCLI {
	baseLogger := (*logger.Logger)(l)
	baseLogger.Fatal(msg, adHocFields...)
	return nestLogger(baseLogger)
}
func (l *LoggerCLI) ErrorFrom(e error, adHocFields ...logger.LogFields) *LoggerCLI {
	baseLogger := (*logger.Logger)(l)
	baseLogger.ErrorFrom(e, adHocFields...)
	return nestLogger(baseLogger)
}
func (l *LoggerCLI) FatalFrom(e error, adHocFields ...logger.LogFields) *LoggerCLI {
	baseLogger := (*logger.Logger)(l)
	baseLogger.FatalFrom(e, adHocFields...)
	return nestLogger(baseLogger)
}

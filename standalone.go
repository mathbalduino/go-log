package golog

// Trace will create a new log with the Trace level, if enabled,
// applying the given adHocFields.
func Trace(msg string, adHocFields ...LogFields) { NewDefault().Trace(msg, adHocFields...) }

// Debug will create a new log with the Debug level, if enabled,
// applying the given adHocFields.
func Debug(msg string, adHocFields ...LogFields) { NewDefault().Debug(msg, adHocFields...) }

// Info will create a new log with the Info level, if enabled,
// applying the given adHocFields.
func Info(msg string, adHocFields ...LogFields) { NewDefault().Info(msg, adHocFields...) }

// Warn will create a new log with the Warn level, if enabled,
// applying the given adHocFields.
func Warn(msg string, adHocFields ...LogFields) { NewDefault().Warn(msg, adHocFields...) }

// Error will create a new log with the Error level, if enabled,
// applying the given adHocFields.
func Error(msg string, adHocFields ...LogFields) { NewDefault().Error(msg, adHocFields...) }

// Fatal will create a new log with the Fatal level, if enabled,
// applying the given adHocFields.
//
// Remember that this method will not call "panic"
func Fatal(msg string, adHocFields ...LogFields) { NewDefault().Fatal(msg, adHocFields...) }

// ErrorFrom will create a new log with the Error level, if enabled,
// using the given error and adHocFields.
func ErrorFrom(e error, adHocFields ...LogFields) { NewDefault().ErrorFrom(e, adHocFields...) }

// FatalFrom will create a new log with the Fatal level, if enabled,
// using the given error and adHocFields.
//
// Remember that this method will not call "panic"
func FatalFrom(e error, adHocFields ...LogFields) { NewDefault().FatalFrom(e, adHocFields...) }

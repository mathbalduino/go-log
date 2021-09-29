package src

// Trace will create a new default log with the Trace level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is variadic just to simulate
// optional params. Latter values will override former ones
func Trace(msg string, adHocFields ...LogFields) { NewDefault().Trace(msg, adHocFields...) }

// Debug will create a new default log with the Debug level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is variadic just to simulate
// optional params. Latter values will override former ones
func Debug(msg string, adHocFields ...LogFields) { NewDefault().Debug(msg, adHocFields...) }

// Info will create a new default log with the Info level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is variadic just to simulate
// optional params. Latter values will override former ones
func Info(msg string, adHocFields ...LogFields) { NewDefault().Info(msg, adHocFields...) }

// Warn will create a new default log with the Warn level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is variadic just to simulate
// optional params. Latter values will override former ones
func Warn(msg string, adHocFields ...LogFields) { NewDefault().Warn(msg, adHocFields...) }

// Error will create a new default log with the Error level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is variadic just to simulate
// optional params. Latter values will override former ones
func Error(msg string, adHocFields ...LogFields) { NewDefault().Error(msg, adHocFields...) }

// ErrorFrom will create a new default log with the Error level, using the
// given error and adHocFields, if enabled.
//
// Note that the 'adHocFields' param is variadic just to simulate
// optional params. Latter values will override former ones
//
// Note that this method will call the default 'ErrorParser'
// function to extract the log message and custom fields from the
// given error. These custom fields will be overridden by the
// 'adHocFields' param
func ErrorFrom(e error, adHocFields ...LogFields) { NewDefault().ErrorFrom(e, adHocFields...) }

// Fatal will create a new default log with the Fatal level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is variadic just to simulate
// optional params. Latter values will override former ones
func Fatal(msg string, adHocFields ...LogFields) { NewDefault().Fatal(msg, adHocFields...) }

// FatalFrom will create a new default log with the Fatal level, using the
// given error and adHocFields, if enabled.
//
// Note that the 'adHocFields' param is variadic just to simulate
// optional params. Latter values will override former ones
//
// Note that this method will call the default 'ErrorParser'
// function to extract the log message and custom fields from the
// given error. These custom fields will be overridden by the
// 'adHocFields' param
func FatalFrom(e error, adHocFields ...LogFields) { NewDefault().FatalFrom(e, adHocFields...) }

package log

type Logger interface {

	// Debugf logs formated message. Consider using Debug in perf-critical places.
	Debugf(format string, arg ...interface{})
	// Infof logs formated message. Consider using Info in perf-critical places.
	Infof(format string, arg ...interface{})
	// Warnf logs formated message. Consider using Warn in perf-critical places.
	Warnf(format string, arg ...interface{})
	// Errorf logs formated message. Consider using Error in perf-critical places.
	Errorf(format string, arg ...interface{})

	// Debug logs msg with debug level.
	Debug(msg string)
	// Info logs msg with info level.
	Info(msg string)
	// Warn logs msg with warn level.
	Warn(msg string)
	// Error logs msg with error level.
	Error(msg string)

	// WithFields adds multiple log message fields.
	WithFields(fields ...Field) Logger
	// WithField adds new log message field. Can be called multiple times, but consider using WithFields.
	WithField(name string, val string) Logger
	// WithSource is a shortcut for WithField("source", val).
	WithSource(source string) Logger
}

type Field struct {
	Key string
	Val string
}

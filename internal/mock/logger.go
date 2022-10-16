package mock

import (
	"sync"

	"github.com/CasinoTrade/CasinoGuest/internal/model/log"
)

type level int

const (
	debug level = iota
	info
	warn
	erro
	fatal
)

type Logger struct {
	m       sync.Mutex
	Records []level
}

// Debugf logs formated message. Consider using Debug in perf-critical places.
func (l *Logger) Debugf(format string, arg ...interface{}) {
	l.m.Lock()
	l.Records = append(l.Records, debug)
	l.m.Unlock()
}

// Infof logs formated message. Consider using Info in perf-critical places.
func (l *Logger) Infof(format string, arg ...interface{}) {
	l.m.Lock()
	l.Records = append(l.Records, info)
	l.m.Unlock()
}

// Warnf logs formated message. Consider using Warn in perf-critical places.
func (l *Logger) Warnf(format string, arg ...interface{}) {
	l.m.Lock()
	l.Records = append(l.Records, warn)
	l.m.Unlock()
}

// Errorf logs formated message. Consider using Error in perf-critical places.
func (l *Logger) Errorf(format string, arg ...interface{}) {
	l.m.Lock()
	l.Records = append(l.Records, debug)
	l.m.Unlock()
}

// Fatalf logs formated message. Consider using Fatal in perf-critical places.
func (l *Logger) Fatalf(format string, arg ...interface{}) {
	l.m.Lock()
	l.Records = append(l.Records, fatal)
	l.m.Unlock()
}

// Debug logs msg with debug level.
func (l *Logger) Debug(msg string) {
	l.m.Lock()
	l.Records = append(l.Records, debug)
	l.m.Unlock()
}

// Info logs msg with info level.
func (l *Logger) Info(msg string) {
	l.m.Lock()
	l.Records = append(l.Records, info)
	l.m.Unlock()
}

// Warn logs msg with warn level.
func (l *Logger) Warn(msg string) {
	l.m.Lock()
	l.Records = append(l.Records, warn)
	l.m.Unlock()
}

// Error logs msg with error level.
func (l *Logger) Error(msg string) {
	l.m.Lock()
	l.Records = append(l.Records, erro)
	l.m.Unlock()
}

// Error logs msg with fatal level and calls os.Exit.
func (l *Logger) Fatal(msg string) {
	l.m.Lock()
	l.Records = append(l.Records, fatal)
	l.m.Unlock()
}

// WithFields adds multiple log message fields.
func (l *Logger) WithFields(fields ...log.Field) log.Logger {
	return l
}

// WithField adds new log message field. Can be called multiple times, but consider using WithFields.
func (l *Logger) WithField(name string, val string) log.Logger {
	return l
}

// WithSource is a shortcut for WithField("source", val).
func (l *Logger) WithSource(source string) log.Logger {
	return l
}

package log

import (
	"time"

	"github.com/CasinoTrade/CasinoGuest/internal/model/config"
	model "github.com/CasinoTrade/CasinoGuest/internal/model/log"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

// Debugf logs formated message. Consider using Debug in perf-critical places.
func (l *Logger) Debugf(format string, arg ...interface{}) {
	l.logger.Debug().Msgf(format, arg...)
}

// Infof logs formated message. Consider using Info in perf-critical places.
func (l *Logger) Infof(format string, arg ...interface{}) {
	l.logger.Info().Msgf(format, arg...)
}

// Warnf logs formated message. Consider using Warn in perf-critical places.
func (l *Logger) Warnf(format string, arg ...interface{}) {
	l.logger.Warn().Msgf(format, arg...)
}

// Errorf logs formated message. Consider using Error in perf-critical places.
func (l *Logger) Errorf(format string, arg ...interface{}) {
	l.logger.Error().Msgf(format, arg...)
}

// Debug logs msg with debug level.
func (l *Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

// Info logs msg with info level.
func (l *Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

// Warn logs msg with warn level.
func (l *Logger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

// Error logs msg with error level.
func (l *Logger) Error(msg string) {
	l.logger.Error().Msg(msg)
}

// WithFields adds multiple log message fields.
func (l *Logger) WithFields(fields ...model.Field) model.Logger {
	lctx := l.logger.With()
	for _, field := range fields {
		lctx = lctx.Str(field.Key, field.Val)
	}
	return &Logger{
		logger: lctx.Logger(),
	}
}

// WithField adds new log message field. Can be called multiple times, but consider using WithFields.
func (l *Logger) WithField(name string, val string) model.Logger {
	return &Logger{
		logger: l.logger.With().Str(name, val).Logger(),
	}
}

// WithSource is a shortcut for WithField("source", val).
func (l *Logger) WithSource(val string) model.Logger {
	return l.WithField("source", val)
}

func New(cfg config.Logger) model.Logger {
	w := zerolog.NewConsoleWriter()
	w.TimeFormat = time.RFC3339
	l := zerolog.New(w).Level(level(cfg.Debug)).With().Timestamp().Logger()
	return &Logger{
		logger: l,
	}
}

func level(debug bool) zerolog.Level {
	if debug {
		return zerolog.DebugLevel
	}
	return zerolog.InfoLevel
}

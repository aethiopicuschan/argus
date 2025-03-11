package argus

import (
	"encoding/json"
	"io"
	"time"

	"github.com/aethiopicuschan/narabi"
)

type Level string

const (
	Info  Level = "INFO"
	Warn  Level = "WARN"
	Error Level = "ERROR"
	Debug Level = "DEBUG"
)

func severity(level Level) int {
	switch level {
	case Debug:
		return 1
	case Info:
		return 2
	case Warn:
		return 3
	case Error:
		return 4
	default:
		return 0
	}
}

// Logger is a logger
type Logger struct {
	writer   io.Writer
	minLevel Level
}

// Option is a functional option for Logger
type Option func(*Logger)

// WithMinLevel is an option to set the minimum level to log
func WithMinLevel(level Level) Option {
	return func(l *Logger) {
		l.minLevel = level
	}
}

// NewLogger creates a new Logger
func NewLogger(writer io.Writer, opts ...Option) (l *Logger) {
	l = &Logger{
		writer:   writer,
		minLevel: Debug,
	}
	for _, opt := range opts {
		opt(l)
	}
	return
}

// Builder is a builder for log
type Builder struct {
	logger *Logger
	level  Level
	om     *narabi.OrderedMap
}

func newBuilder(logger *Logger, level Level) (b *Builder) {
	b = &Builder{
		logger: logger,
		level:  level,
		om:     narabi.New(),
	}
	b.Add("level", level)
	b.Add("time", time.Now())
	return
}

// Info is a method to create a log with level INFO
func (l *Logger) Info() *Builder {
	return newBuilder(l, Info)
}

// Warn is a method to create a log with level WARN
func (l *Logger) Warn() *Builder {
	return newBuilder(l, Warn)
}

// Error is a method to create a log with level ERROR
func (l *Logger) Error() *Builder {
	return newBuilder(l, Error)
}

// Debug is a method to create a log with level DEBUG
func (l *Logger) Debug() *Builder {
	return newBuilder(l, Debug)
}

// Add is a method to add key-value pair to the log
func (b *Builder) Add(key string, value interface{}) *Builder {
	b.om.Set(key, value)
	return b
}

// Remove is a method to remove key-value pair from the log
func (b *Builder) Remove(key string) *Builder {
	b.om.Delete(key)
	return b
}

// Print is a method to print the log
func (b *Builder) Print() (err error) {
	if severity(b.level) < severity(b.logger.minLevel) {
		return
	}
	jsonData, err := json.Marshal(b.om)
	if err != nil {
		return
	}
	jsonData = append(jsonData, '\n')
	_, err = b.logger.writer.Write(jsonData)
	return
}

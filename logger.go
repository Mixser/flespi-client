package flespi

import (
	"fmt"
	"log"
	"os"
)

// Logger is an interface for logging client operations
type Logger interface {
	// Debugf logs a debug message
	Debugf(format string, args ...interface{})

	// Infof logs an info message
	Infof(format string, args ...interface{})

	// Warnf logs a warning message
	Warnf(format string, args ...interface{})

	// Errorf logs an error message
	Errorf(format string, args ...interface{})
}

// WithLogger sets a custom logger for the client
func WithLogger(logger Logger) ClientOption {
	return func(c *Client) {
		c.Logger = logger
	}
}

// NoOpLogger is a logger that does nothing
type NoOpLogger struct{}

func (l *NoOpLogger) Debugf(format string, args ...interface{}) {}
func (l *NoOpLogger) Infof(format string, args ...interface{})  {}
func (l *NoOpLogger) Warnf(format string, args ...interface{})  {}
func (l *NoOpLogger) Errorf(format string, args ...interface{}) {}

// StdLogger is a simple logger that writes to standard output
type StdLogger struct {
	logger *log.Logger
	level  LogLevel
}

// LogLevel represents the logging level
type LogLevel int

const (
	// LogLevelDebug enables all logs including debug messages
	LogLevelDebug LogLevel = iota
	// LogLevelInfo enables info, warning, and error logs
	LogLevelInfo
	// LogLevelWarn enables warning and error logs
	LogLevelWarn
	// LogLevelError enables only error logs
	LogLevelError
)

// NewStdLogger creates a new standard logger with the specified log level
func NewStdLogger(level LogLevel) *StdLogger {
	return &StdLogger{
		logger: log.New(os.Stdout, "[flespi] ", log.LstdFlags),
		level:  level,
	}
}

func (l *StdLogger) Debugf(format string, args ...interface{}) {
	if l.level <= LogLevelDebug {
		l.logger.Printf("[DEBUG] "+format, args...)
	}
}

func (l *StdLogger) Infof(format string, args ...interface{}) {
	if l.level <= LogLevelInfo {
		l.logger.Printf("[INFO] "+format, args...)
	}
}

func (l *StdLogger) Warnf(format string, args ...interface{}) {
	if l.level <= LogLevelWarn {
		l.logger.Printf("[WARN] "+format, args...)
	}
}

func (l *StdLogger) Errorf(format string, args ...interface{}) {
	if l.level <= LogLevelError {
		l.logger.Printf("[ERROR] "+format, args...)
	}
}

// logRequest logs an outgoing HTTP request
func (c *Client) logRequest(method, endpoint string, payload interface{}) {
	if c.Logger == nil {
		return
	}

	if payload != nil {
		c.Logger.Debugf("Request: %s %s with payload", method, endpoint)
	} else {
		c.Logger.Debugf("Request: %s %s", method, endpoint)
	}
}

// logResponse logs an HTTP response
func (c *Client) logResponse(method, endpoint string, statusCode int, err error) {
	if c.Logger == nil {
		return
	}

	if err != nil {
		c.Logger.Errorf("Response: %s %s failed - %v", method, endpoint, err)
	} else {
		c.Logger.Debugf("Response: %s %s succeeded (status %d)", method, endpoint, statusCode)
	}
}

// logError is a helper to log errors
func (c *Client) logError(format string, args ...interface{}) {
	if c.Logger != nil {
		c.Logger.Errorf(format, args...)
	}
}

// logDebug is a helper to log debug messages
func (c *Client) logDebug(format string, args ...interface{}) {
	if c.Logger != nil {
		c.Logger.Debugf(format, args...)
	}
}

// EnableDebugLogging is a convenience function to enable debug logging to stdout
func EnableDebugLogging() ClientOption {
	return WithLogger(NewStdLogger(LogLevelDebug))
}

// EnableInfoLogging is a convenience function to enable info logging to stdout
func EnableInfoLogging() ClientOption {
	return WithLogger(NewStdLogger(LogLevelInfo))
}

// String returns a string representation of a log level
func (l LogLevel) String() string {
	switch l {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	default:
		return fmt.Sprintf("UNKNOWN(%d)", l)
	}
}

package flespi

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestNoOpLogger(t *testing.T) {
	logger := &NoOpLogger{}

	// Should not panic
	logger.Debugf("test")
	logger.Infof("test")
	logger.Warnf("test")
	logger.Errorf("test")
}

func TestStdLogger_LogLevels(t *testing.T) {
	tests := []struct {
		name     string
		level    LogLevel
		logFunc  func(*StdLogger)
		expected bool
	}{
		{
			name:  "Debug level logs debug",
			level: LogLevelDebug,
			logFunc: func(l *StdLogger) {
				l.Debugf("test")
			},
			expected: true,
		},
		{
			name:  "Info level does not log debug",
			level: LogLevelInfo,
			logFunc: func(l *StdLogger) {
				l.Debugf("test")
			},
			expected: false,
		},
		{
			name:  "Info level logs info",
			level: LogLevelInfo,
			logFunc: func(l *StdLogger) {
				l.Infof("test")
			},
			expected: true,
		},
		{
			name:  "Warn level does not log info",
			level: LogLevelWarn,
			logFunc: func(l *StdLogger) {
				l.Infof("test")
			},
			expected: false,
		},
		{
			name:  "Error level only logs errors",
			level: LogLevelError,
			logFunc: func(l *StdLogger) {
				l.Warnf("test")
			},
			expected: false,
		},
		{
			name:  "Error level logs errors",
			level: LogLevelError,
			logFunc: func(l *StdLogger) {
				l.Errorf("test")
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			logger := &StdLogger{
				logger: log.New(&buf, "", 0),
				level:  tt.level,
			}

			tt.logFunc(logger)

			output := buf.String()
			hasOutput := len(output) > 0

			if hasOutput != tt.expected {
				t.Errorf("Expected output=%v, got output=%v (output: %s)", tt.expected, hasOutput, output)
			}
		})
	}
}

func TestLogLevelString(t *testing.T) {
	tests := []struct {
		level    LogLevel
		expected string
	}{
		{LogLevelDebug, "DEBUG"},
		{LogLevelInfo, "INFO"},
		{LogLevelWarn, "WARN"},
		{LogLevelError, "ERROR"},
		{LogLevel(99), "UNKNOWN(99)"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.level.String()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestClient_WithLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := &StdLogger{
		logger: log.New(&buf, "", 0),
		level:  LogLevelDebug,
	}

	client, _ := NewClient("https://flespi.io", "test-token", WithLogger(logger))

	if client.Logger == nil {
		t.Errorf("Expected logger to be set, got nil")
	}

	// Test that logging works
	client.logDebug("test message")

	output := buf.String()
	if !strings.Contains(output, "test message") {
		t.Errorf("Expected log output to contain 'test message', got: %s", output)
	}
}

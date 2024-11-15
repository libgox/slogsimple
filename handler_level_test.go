package slogsimple

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slog"
)

func TestLogLevelFiltering(t *testing.T) {
	buffer := &bytes.Buffer{}
	handler := NewHandler(&Config{
		Output:   buffer,
		MinLevel: slog.LevelWarn,
	})
	logger := slog.New(handler)

	// INFO level log should not appear
	logger.InfoContext(context.Background(), "This should not be logged")
	assert.Equal(t, 0, buffer.Len(), "Expected no log output for INFO level")

	// WARN level log should appear
	logger.WarnContext(context.Background(), "This is a warning")
	expected := "This is a warning"
	assert.Contains(t, buffer.String(), expected, "Expected log to contain warning message")

	// ERROR level log should also appear
	logger.ErrorContext(context.Background(), "This is an error")
	expectedError := "This is an error"
	assert.Contains(t, buffer.String(), expectedError, "Expected log to contain error message")
}

func TestGroupLogLevelFiltering(t *testing.T) {
	buffer := &bytes.Buffer{}
	handler := NewHandler(&Config{
		Output:   buffer,
		MinLevel: slog.LevelInfo,
		GroupLevels: map[string]slog.Level{
			"auth": slog.LevelDebug,
			"db":   slog.LevelError,
		},
	})
	logger := slog.New(handler)

	logger.WithGroup("auth").DebugContext(context.Background(), "Auth debug log")
	assert.Contains(t, buffer.String(), "Auth debug log", "Expected auth debug log to appear")

	logger.WithGroup("auth").InfoContext(context.Background(), "Auth info log")
	assert.Contains(t, buffer.String(), "Auth info log", "Expected auth info log to appear")

	buffer.Reset()

	logger.WithGroup("db").WarnContext(context.Background(), "DB warning log")
	assert.NotContains(t, buffer.String(), "DB warning log", "Expected db warning log to be filtered")

	logger.WithGroup("db").ErrorContext(context.Background(), "DB error log")
	assert.Contains(t, buffer.String(), "DB error log", "Expected db error log to appear")

	buffer.Reset()

	logger.WithGroup("other").InfoContext(context.Background(), "Other info log")
	assert.Contains(t, buffer.String(), "Other info log", "Expected other info log to appear")

	logger.WithGroup("other").DebugContext(context.Background(), "Other debug log")
	assert.NotContains(t, buffer.String(), "Other debug log", "Expected other debug log to be filtered")
}

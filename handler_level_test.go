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
	handler := NewHandler(buffer, slog.LevelWarn)
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

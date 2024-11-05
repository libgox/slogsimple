package slogsimple

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slog"
)

type ErrorWriter struct{}

func (w *ErrorWriter) Write(p []byte) (int, error) {
	return 0, errors.New("write error")
}

func TestHandleError(t *testing.T) {
	errorWriter := &ErrorWriter{}
	handler := NewHandler(&Config{
		Output:   errorWriter,
		MinLevel: slog.LevelError,
	})
	logger := slog.New(handler)

	logger.ErrorContext(context.Background(), "This is an error message", slog.String("key", "value"))

	buffer := &bytes.Buffer{}
	handler = NewHandler(&Config{
		Output:   buffer,
		MinLevel: slog.LevelError,
	})
	logger = slog.New(handler)

	logger.ErrorContext(context.Background(), "This is a second error message", slog.String("key", "value"))
	assert.Contains(t, buffer.String(), "ERROR This is a second error message key=value", "Expected error message to be logged to buffer")
}

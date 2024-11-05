package slogsimple

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slog"
)

func TestLogFormatBasic(t *testing.T) {
	buffer := &bytes.Buffer{}
	handler := NewHandler(&Config{
		Output:   buffer,
		MinLevel: slog.LevelInfo,
	})
	logger := slog.New(handler)

	logger.InfoContext(context.Background(), "Basic log message", slog.String("key", "value"))
	timestamp := time.Now().Format(time.RFC3339)
	expected := timestamp + " INFO Basic log message key=value "
	assert.Contains(t, buffer.String(), expected, "Expected log to contain basic formatted log with level")
}

func TestLogFormatWithGroup(t *testing.T) {
	buffer := &bytes.Buffer{}
	handler := NewHandler(&Config{
		Output:   buffer,
		MinLevel: slog.LevelInfo,
	}).WithGroup("module")
	logger := slog.New(handler)

	logger.InfoContext(context.Background(), "Grouped log message", slog.String("key", "value"))
	expected := "INFO Grouped log message module.key=value"
	assert.Contains(t, buffer.String(), expected, "Expected log to contain grouped key=value with level")
}

func TestLogFormatWithAttrs(t *testing.T) {
	buffer := &bytes.Buffer{}
	handler := NewHandler(&Config{
		Output:   buffer,
		MinLevel: slog.LevelInfo,
	}).WithAttrs([]slog.Attr{
		slog.String("app", "myApp"),
		slog.String("env", "production"),
	})
	logger := slog.New(handler)

	logger.InfoContext(context.Background(), "Log with attributes", slog.String("key", "value"))
	expected := "INFO Log with attributes app=myApp env=production key=value"
	assert.Contains(t, buffer.String(), expected, "Expected log to contain attributes with key=value and level")
}

package slogsimple

import (
	"context"
	"io"
	"strings"
	"time"

	"golang.org/x/exp/slog"
)

type Config struct {
	Output   io.Writer
	MinLevel slog.Level
}

type Handler struct {
	output   io.Writer
	minLevel slog.Level

	groupName string
	attrs     []slog.Attr
}

func NewHandler(config *Config) slog.Handler {
	return &Handler{output: config.Output, minLevel: config.MinLevel}
}

func (h *Handler) Enabled(context context.Context, level slog.Level) bool {
	return level >= h.minLevel
}

func (h *Handler) Handle(context context.Context, r slog.Record) error {
	timestamp := r.Time.Format(time.RFC3339)
	message := r.Message

	var builder strings.Builder
	builder.WriteString(timestamp)
	builder.WriteString(" ")
	builder.WriteString(message)
	builder.WriteString(" ")

	for _, attr := range h.attrs {
		builder.WriteString(attr.Key)
		builder.WriteString("=")
		builder.WriteString(attr.Value.String())
		builder.WriteString(" ")
	}

	r.Attrs(func(attr slog.Attr) bool {
		if h.groupName != "" {
			builder.WriteString(h.groupName)
			builder.WriteString(".")
		}
		builder.WriteString(attr.Key)
		builder.WriteString("=")
		builder.WriteString(attr.Value.String())
		builder.WriteString(" ")
		return true
	})

	builder.WriteString("\n")
	_, err := h.output.Write([]byte(builder.String()))
	return err
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newAttrs := append(h.attrs, attrs...)
	return &Handler{groupName: h.groupName, attrs: newAttrs, output: h.output}
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{groupName: name, attrs: h.attrs, output: h.output}
}

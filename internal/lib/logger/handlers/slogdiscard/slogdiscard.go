package slogdiscard

import (
	"context"
	"log/slog"
)

func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHandler())
}

type DiscardHadler struct{}

func NewDiscardHandler() *DiscardHadler {
	return &DiscardHadler{}
}

func (h *DiscardHadler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}

func (h *DiscardHadler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (h *DiscardHadler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

func (h *DiscardHadler) WithGroup(name string) slog.Handler {
	return h
}

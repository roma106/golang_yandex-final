package logger

import "log/slog"

func Info(message string) {
	slog.Info(message)
}

func Error(message string) {
	slog.Error(message)
}

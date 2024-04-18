package logger

import (
	"fmt"
	"log/slog"
)

func Info(message string) {
	slog.Info(fmt.Sprintf("| %s", message))
}

func Error(message string) {
	slog.Error(fmt.Sprintf("| %s", message))
}

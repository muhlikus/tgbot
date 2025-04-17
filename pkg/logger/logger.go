package logger

import (
	"io"
	"log/slog"
	"os"
)

func NewFileLogger(filePath string) (*slog.Logger, *slog.LevelVar) {

	var logWriter io.Writer

	logWriter, err := os.Create(filePath)
	if err != nil {
		logWriter = os.Stdout
	}

	logLeveler := &slog.LevelVar{}
	logHandler := slog.NewTextHandler(logWriter, &slog.HandlerOptions{Level: logLeveler})
	return slog.New(logHandler), logLeveler
}

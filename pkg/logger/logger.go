package logger

import (
	"io"
	"log"
	"log/slog"
	"os"
)

func NewFileLogger() (*slog.Logger, *slog.LevelVar) {
	var writer io.Writer

	filePath := os.Getenv("LOG_PATH")
	if filePath == "" {
		log.Fatal("Environment LOG_PATH must be set")
	}

	writer, err := os.Create(filePath)
	if err != nil {
		writer = os.Stdout
	}

	leveler := &slog.LevelVar{}
	handler := slog.NewJSONHandler(writer, &slog.HandlerOptions{Level: leveler})
	return slog.New(handler), leveler
}

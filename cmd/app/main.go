package main

import (
	"log/slog"
	"os"

	repository "github.com/muhlikus/tgbot/internal/repository/external/telegram"
	"github.com/muhlikus/tgbot/pkg/logger"

	"github.com/caarlos0/env/v11"
	"github.com/muhlikus/telegramclient"
)

const defaultLogFileName = "tgbot.log"

func main() {
	var cfg config

	logFilePath := os.Getenv("LOG_PATH")
	if logFilePath == "" {
		logFilePath = defaultLogFileName
	}

	logger, leveler := logger.NewFileLogger(logFilePath)
	leveler.Set(slog.LevelDebug)
	//slog.SetDefault(logger)

	err := env.Parse(&cfg)
	if err != nil {
		logger.Error("parsing config", slog.Any("error", err))
		return
	}

	telegramClient, err := telegramclient.New(cfg.TgClient)
	if err != nil {
		logger.Error("creating telegram client", slog.Any("error", err))
		return
	}

	rep, err := repository.NewRepository(telegramClient)
	if err != nil {
		logger.Error("creating repository", slog.Any("error", err))
		return
	}
	_ = rep
}

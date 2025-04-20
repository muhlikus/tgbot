package main

import (
	"log"
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

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("failed to parse config %v", err)
	}

	logger, leveler := logger.NewFileLogger()
	leveler.Set(slog.LevelDebug)
	//slog.SetDefault(logger)

	_, cfg.debug = os.LookupEnv("DEBUG")
	if cfg.debug {
		leveler.Set(slog.LevelDebug)
	}

	telegramClient, err := telegramclient.New(cfg.TgClient)
	if err != nil {
		logger.Error("failed to create telegram client", slog.Any("error", err))
		return
	}

	rep, err := repository.NewRepository(telegramClient)
	if err != nil {
		logger.Error("failed to create repository", slog.Any("error", err))
		return
	}
	_ = rep
}

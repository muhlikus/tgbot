package main

import (
	"log"
	"os"

	repository "github.com/muhlikus/tgbot/internal/repository/external/telegram"

	"github.com/muhlikus/telegramclient"
)

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatalln("TELEGRAM_BOT_TOKEN environment must be set")
	}

	config := &Config{Token: botToken}

	telegramClient, err := telegramclient.New(telegramclient.Config{Token: config.Token})
	if err != nil {
		log.Fatalf("Failed to create telegram client: %v", err)
	}

	rep, err := repository.NewRepository(telegramClient)
	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}
	_ = rep
}

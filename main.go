package main

import (
	"os"

	"github.com/muhlikus/telegramclient"
)

func main() {

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		panic("TELEGRAM_BOT_TOKEN environment must be set")
	}

	// Create a new TelegramClient instance
	_, _ = telegramclient.New(
		telegramclient.Config{
			Token: botToken,
		})
}

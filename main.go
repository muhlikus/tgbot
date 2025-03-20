package main

import "github.com/muhlikus/telegramclient"

func main() {
	// Create a new TelegramClient instance

	_, _ = telegramclient.New(
		telegramclient.Config{
			Token: "fake-token",
		})
}

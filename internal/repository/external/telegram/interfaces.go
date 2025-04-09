package repository

import (
	"bytes"

	"github.com/muhlikus/telegramclient"
)

type telegramClient interface {
	GetUpdates() ([]telegramclient.Update, error)
	SendMessage(chatID int, text string) (*telegramclient.Message, error)
	SendDocument(chatID int, fileName string, fileBuff *bytes.Buffer) (*telegramclient.Message, error)
}

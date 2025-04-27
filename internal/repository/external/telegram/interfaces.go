package repository

import (
	"bytes"

	"github.com/muhlikus/telegramclient"
)

//go:generate mockgen -source=interfaces.go -destination=./mock/interfaces.go -package=mock
type telegramClient interface {
	GetUpdates() ([]telegramclient.Update, error)
	SendMessage(chatID int, text string) (*telegramclient.Message, error)
	SendDocument(chatID int, fileName string, fileBuff *bytes.Buffer) (*telegramclient.Message, error)
}

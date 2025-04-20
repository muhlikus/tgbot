package service

import (
	"bytes"

	"github.com/muhlikus/tgbot/internal/models"
)

type repository interface {
	GetMessages() ([]models.Message, error)
	SendMessage(chatID int, text string) error
	SendFile(chatID int, fileName string, fileBuff *bytes.Buffer) error
}

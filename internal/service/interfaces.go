package service

import (
	"bytes"

	"github.com/muhlikus/tgbot/internal/models"
)

//go:generate mockgen -source=interfaces.go -destination=./mock/interfaces.go -package=mock
type repository interface {
	GetMessages() ([]models.Message, error)
	SendMessage(chatID int, text string) error
	SendFile(chatID int, fileName string, fileBuff *bytes.Buffer) error
}

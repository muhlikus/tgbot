package repository

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/muhlikus/telegramclient"
)

type Repository struct {
	client telegramClient
}

func NewRepository(t telegramClient) (*Repository, error) {
	if t == nil {
		return nil, errors.New("telegram client is nil")
	}

	return &Repository{client: t}, nil
}

func (r *Repository) GetUpdates() ([]telegramclient.Update, error) {
	return r.client.GetUpdates()
}

// TODO надо подумать, стоит ли тут возвращать сообщение
func (r *Repository) SendMessage(chatID int, text string) error {
	_, err := r.client.SendMessage(chatID, text)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

// TODO надо подумать, стоит ли тут возвращать сообщение
func (r *Repository) SendDocument(chatID int, fileName string, fileBuff *bytes.Buffer) error {
	_, err := r.client.SendDocument(chatID, fileName, fileBuff)
	if err != nil {
		return fmt.Errorf("failed to send document: %w", err)
	}

	return nil
}

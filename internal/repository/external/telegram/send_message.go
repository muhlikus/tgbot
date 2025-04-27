package repository

import (
	"fmt"
)

// TODO надо подумать, стоит ли тут возвращать сообщение
func (r *Repository) SendMessage(chatID int, text string) error {
	_, err := r.client.SendMessage(chatID, text)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

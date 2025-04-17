package repository

import (
	"bytes"
	"fmt"
)

// TODO надо подумать, стоит ли тут возвращать сообщение
func (r *Repository) SendDocument(chatID int, fileName string, fileBuff *bytes.Buffer) error {
	_, err := r.client.SendDocument(chatID, fileName, fileBuff)
	if err != nil {
		return fmt.Errorf("failed to send document: %w", err)
	}

	return nil
}

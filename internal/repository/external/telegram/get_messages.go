package repository

import (
	"fmt"

	"github.com/muhlikus/tgbot/internal/models"
)

func (r *Repository) GetMessages() ([]models.Message, error) {
	updates, err := r.client.GetUpdates()
	if err != nil {
		return nil, fmt.Errorf("failed to get messages: %w", err)
	}

	messages := make([]models.Message, 0, len(updates))

	for _, v := range updates {
		if v.Message == nil {
			continue
		}

		message := models.Message{
			Id:     v.Message.MessageId,
			Date:   v.Message.Date,
			ChatID: int(v.Message.Chat.Id),
		}

		switch {
		case v.Message.Text != "":
			message.Text = v.Message.Text
			message.Type = models.TextMessage
		case v.Message.Document != nil:
			message.File = &models.File{
				ID:       v.Message.Document.FileID,
				UniqueID: v.Message.Document.FileUniqueID,
				Size:     v.Message.Document.FileSize,
				Name:     v.Message.Document.FileName,
				MimeType: v.Message.Document.MimeType,
			}
		}

		messages = append(messages, message)
	}

	return messages, nil
}

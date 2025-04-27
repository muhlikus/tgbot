package service

import (
	"bytes"
	"errors"

	"github.com/muhlikus/tgbot/internal/models"
)

type Service struct {
	repo repository
}

func NewService(r repository) (*Service, error) {
	if r == nil {
		return nil, errors.New("repository is nil")
	}

	return &Service{repo: r}, nil
}

func (s *Service) GetMessages() ([]models.Message, error) {
	return s.repo.GetMessages()
}

func (s *Service) SendMessage(chatID int, text string) error {
	return s.repo.SendMessage(chatID, text)
}

func (s *Service) SendFile(chatID int, fileName string, fileBuff *bytes.Buffer) error {
	return s.repo.SendFile(chatID, fileName, fileBuff)
}

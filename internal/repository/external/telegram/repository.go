package repository

import (
	"errors"
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

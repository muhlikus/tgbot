package service

import "errors"

type Service struct {
	Repository repository
}

func NewService(r repository) (*Service, error) {
	if r == nil {
		return nil, errors.New("repository is nil")
	}

	return &Service{}, nil
}

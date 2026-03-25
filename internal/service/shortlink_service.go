package service

import (
	"errors"
	"net/url"
	"time"
	"urlshortener/internal/model"
	"urlshortener/internal/repository"
	"urlshortener/internal/utils"
)

var (
	ErrInvalidURL = errors.New("invalid original url format")
)

type shortlinkService struct {
	repo repository.ShortlinkRepository
}

func NewShortlinkService(repo repository.ShortlinkRepository) ShortlinkService {
	return &shortlinkService{repo: repo}
}

func (s *shortlinkService) CreateShortlink(originalURL string) (*model.Shortlink, error) {
	// Validate URL
	if !isValidURL(originalURL) {
		return nil, ErrInvalidURL
	}

	code, err := utils.GenerateShortCode()

	shortlink := &model.Shortlink{
		ID:          code,
		OriginalURL: originalURL,
		CreatedAt:   time.Now().UTC(),
	}

	err = s.repo.Create(shortlink)
	if err != nil {
		return nil, err
	}

	return shortlink, nil
}

func isValidURL(str string) bool {
	u, err := url.ParseRequestURI(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

package service

import "urlshortener/internal/model"

type ShortlinkService interface {
	CreateShortlink(originalURL string) (*model.Shortlink, error)
	GetShortlink(id string) (*model.Shortlink, error)
}

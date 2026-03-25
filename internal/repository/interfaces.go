package repository

import "urlshortener/internal/model"

type ShortlinkRepository interface {
	Create(shortlink *model.Shortlink) error
	FindByID(id string) (*model.Shortlink, error)
	FindByOriginalURL(originalURL string) (*model.Shortlink, error)
}

package repository

import "urlshortener/internal/model"

type ShortlinkRepository interface {
	Create(shortlink *model.Shortlink) error
}

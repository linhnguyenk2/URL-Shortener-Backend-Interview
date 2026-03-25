package repository

import (
	"errors"
	"urlshortener/internal/model"

	"gorm.io/gorm"
)

var ErrNotFound = errors.New("shortlink not found")

type ShortlinkRepo struct {
	db *gorm.DB
}

func NewShortlinkRepository(db *gorm.DB) ShortlinkRepository {
	return &ShortlinkRepo{db: db}
}

func (r *ShortlinkRepo) Create(shortlink *model.Shortlink) error {
	return r.db.Create(shortlink).Error
}

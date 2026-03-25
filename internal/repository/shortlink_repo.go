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

func (r *ShortlinkRepo) FindByID(id string) (*model.Shortlink, error) {
	var shortlink model.Shortlink
	result := r.db.Where("id = ?", id).Limit(1).Find(&shortlink)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, ErrNotFound
	}
	return &shortlink, nil
}

func (r *ShortlinkRepo) FindByOriginalURL(originalURL string) (*model.Shortlink, error) {
	var shortlink model.Shortlink
	result := r.db.Where("original_url = ?", originalURL).Limit(1).Find(&shortlink)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, ErrNotFound
	}
	return &shortlink, nil
}

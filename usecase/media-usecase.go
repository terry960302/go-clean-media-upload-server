package usecase

import (
	"gorm.io/gorm"
)

// handle business logic of 'media'
type MediaServiceImpl struct {
	DB *gorm.DB
}

type MediaService interface {
	GetAllMedia()
}

func NewMediaService(db *gorm.DB) *MediaServiceImpl {
	return &MediaServiceImpl{
		DB: db,
	}
}

func (m *MediaServiceImpl) GetAllMedia() error {
	// db := m.DB

	return nil

}

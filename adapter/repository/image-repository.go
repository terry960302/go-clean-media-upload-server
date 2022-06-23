package repository

import (
	"media-upload-server/domain"

	"gorm.io/gorm"
)

type ImageMetadataRepository struct {
	domain.ImageMetadataRepository
	DB *gorm.DB
}

func NewImageMetadataRepository(db *gorm.DB) *ImageMetadataRepository {
	repo := &ImageMetadataRepository{DB: db}
	repo.ImageMetadataRepository = interface{}(repo).(*ImageMetadataRepository)
	return repo
}

func (ir *ImageMetadataRepository) GetAll() ([]domain.ImageMetadata, error) {
	return []domain.ImageMetadata{}, nil
}

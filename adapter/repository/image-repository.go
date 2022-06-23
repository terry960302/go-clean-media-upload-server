package repository

import (
	"media-upload-server/domain"

	"gorm.io/gorm"
)

type ImageMetadataRepository struct {
	DB *gorm.DB
}

func NewImageMetadataRepository(db *gorm.DB) *ImageMetadataRepository {
	return &ImageMetadataRepository{DB: db}
}

func (ir *ImageMetadataRepository) GetAll() ([]domain.ImageMetadata, error) {
	return []domain.ImageMetadata{}, nil
}

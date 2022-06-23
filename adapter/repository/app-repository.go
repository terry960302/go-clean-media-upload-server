package repository

import "gorm.io/gorm"

type AppRepository struct {
	ImageRepo ImageMetadataRepository
	MediaRepo MediaMetadataRepository
	// VideoRepo VideoMetadataRepository
}

func NewAppRepository(db *gorm.DB) *AppRepository {
	return &AppRepository{
		ImageRepo: *NewImageMetadataRepository(db),
		MediaRepo: *NewMediaMetadataRepository(db),
	}
}

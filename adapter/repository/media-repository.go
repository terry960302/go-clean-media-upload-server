package repository

import (
	"fmt"
	"log"
	"media-upload-server/domain"

	"gorm.io/gorm"
)

type MediaMetadataRepository struct {
	domain.MediaMetadataRepository
	DB *gorm.DB
}

func NewMediaMetadataRepository(db *gorm.DB) *MediaMetadataRepository {
	repo := &MediaMetadataRepository{
		DB: db,
	}
	repo.MediaMetadataRepository = interface{}(repo).(*MediaMetadataRepository) // for 'override'
	return repo
}

func (m *MediaMetadataRepository) GetAll() ([]domain.MediaMetadata, error) {

	media := []domain.MediaMetadata{}

	result := m.DB.Find(&media)
	fmt.Printf("%s data loaded", fmt.Sprint(result.RowsAffected))

	if err := result.Error; err != nil {
		log.Fatal(err)
		return nil, err
	}

	return media, nil
}

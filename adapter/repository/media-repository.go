package repository

import (
	"fmt"
	"log"

	"github.com/terry960302/go-clean-media-upload-server/domain"

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

func (m *MediaMetadataRepository) Create(media domain.MediaMetadata) (uint, error) {
	trx := m.DB.Create(&media)

	err := trx.Error
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	fmt.Printf("%s MediaImage created", fmt.Sprint(trx.RowsAffected))
	return uint(media.ID), nil
}

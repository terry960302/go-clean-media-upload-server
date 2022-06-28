package repository

import (
	"fmt"
	"log"

	"github.com/terry960302/go-clean-media-upload-server/domain"
	"gorm.io/gorm"
)

type VideoMetadataRepository struct {
	domain.VideoMetadataRepository
	DB *gorm.DB
}

func NewVideoMetadataRepository(db *gorm.DB) *VideoMetadataRepository {
	repo := &VideoMetadataRepository{DB: db}
	repo.VideoMetadataRepository = interface{}(repo).(*VideoMetadataRepository)
	return repo
}

func (i *VideoMetadataRepository) Create(video domain.VideoMetadata) (uint, error) {
	trx := i.DB.Create(&video)
	err := trx.Error
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	fmt.Printf("%s ImageMetadata is created\n", fmt.Sprint(trx.RowsAffected))
	return uint(video.ID), nil
}

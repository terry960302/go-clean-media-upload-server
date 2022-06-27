package repository

import (
	"fmt"
	"log"

	"github.com/terry960302/go-clean-media-upload-server/domain"

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

func (i *ImageMetadataRepository) GetAll() ([]domain.ImageMetadata, error) {
	return []domain.ImageMetadata{}, nil
}

func (i *ImageMetadataRepository) Create(image domain.ImageMetadata) (uint, error) {
	trx := i.DB.Create(&image)
	err := trx.Error
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	fmt.Printf("%s ImageMetadata is created", fmt.Sprint(trx.RowsAffected))
	return uint(image.ID), nil
}

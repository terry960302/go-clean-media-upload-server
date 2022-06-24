package usecase

import (
	"media-upload-server/adapter/repository"
	"media-upload-server/domain"
)

type MediaMetadataUsecase struct {
	domain.MediaMetadataUsecase
	mediaRepo repository.MediaMetadataRepository
}

func NewMediaMetadataUsecase(mediaRepo repository.MediaMetadataRepository) *MediaMetadataUsecase {
	usecase := &MediaMetadataUsecase{mediaRepo: mediaRepo}
	usecase.MediaMetadataUsecase = interface{}(usecase).(*MediaMetadataUsecase)
	return usecase
}

func (m *MediaMetadataUsecase) GetAllMedia() ([]domain.MediaMetadata, error) {
	return nil, nil
}

func (m *MediaMetadataUsecase) CreateMediaImage() error {
	return nil
}

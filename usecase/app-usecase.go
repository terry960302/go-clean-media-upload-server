package usecase

import "github.com/terry960302/go-clean-media-upload-server/adapter/repository"

type AppUsecase struct {
	ImageUsecase ImageMetadataUsecase
	MediaUsecase MediaMetadataUsecase
}

func NewAppUsecase(imageRepo repository.ImageMetadataRepository, mediaRepo repository.MediaMetadataRepository) *AppUsecase {
	return &AppUsecase{
		ImageUsecase: *NewImageUsecase(imageRepo, mediaRepo),
		MediaUsecase: *NewMediaMetadataUsecase(mediaRepo),
	}
}

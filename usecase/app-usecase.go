package usecase

import "github.com/terry960302/go-clean-media-upload-server/adapter/repository"

type AppUsecase struct {
	ImageUsecase ImageMetadataUsecase
	MediaUsecase MediaMetadataUsecase
	VideoUsecase VideoMetadataUsecase
}

func NewAppUsecase(appRepo *repository.AppRepository) *AppUsecase {
	imgUsecase := *NewImageMetadataUsecase(appRepo.ImageRepo, appRepo.MediaRepo)
	videoUsecase := *NewVideoMetadataUsecase(appRepo.VideoRepo, appRepo.MediaRepo)
	return &AppUsecase{
		ImageUsecase: *NewImageMetadataUsecase(appRepo.ImageRepo, appRepo.MediaRepo),
		MediaUsecase: *NewMediaMetadataUsecase(appRepo.MediaRepo, imgUsecase, videoUsecase)}
}

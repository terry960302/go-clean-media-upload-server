package controller

import "github.com/terry960302/go-clean-media-upload-server/usecase"

type AppController struct {
	ImageCtrl ImageMetadataController
	MediaCtrl MediaMetadataController
}

func NewAppController(imageUsecase usecase.ImageMetadataUsecase, mediaUsecase usecase.MediaMetadataUsecase) *AppController {
	return &AppController{
		ImageCtrl: *NewImageMetadataController(imageUsecase),
		MediaCtrl: *NewMediaMetadataController(mediaUsecase),
	}
}

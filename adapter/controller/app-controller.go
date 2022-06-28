package controller

import "github.com/terry960302/go-clean-media-upload-server/usecase"

type AppController struct {
	ImageCtrl ImageMetadataController
	MediaCtrl MediaMetadataController
}

func NewAppController(appUsecase *usecase.AppUsecase) *AppController {
	return &AppController{
		ImageCtrl: *NewImageMetadataController(appUsecase.ImageUsecase),
		MediaCtrl: *NewMediaMetadataController(appUsecase.MediaUsecase),
	}
}

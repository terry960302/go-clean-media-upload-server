package controller

import (
	"media-upload-server/domain"
	"media-upload-server/usecase"
)

type MediaMetadataController struct {
	domain.MediaMetadataController
	mediaUsecase usecase.MediaMetadataUsecase
}

func NewMediaMetadataController(mediaUsecase usecase.MediaMetadataUsecase) *MediaMetadataController {
	ctrl := &MediaMetadataController{mediaUsecase: mediaUsecase}
	ctrl.MediaMetadataController = interface{}(ctrl).(*MediaMetadataController)
	return ctrl
}

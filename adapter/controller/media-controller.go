package controller

import (
	"github.com/terry960302/go-clean-media-upload-server/domain"
	"github.com/terry960302/go-clean-media-upload-server/usecase"
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

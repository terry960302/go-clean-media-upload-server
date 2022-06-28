package controller

import (
	"github.com/terry960302/go-clean-media-upload-server/domain"
	"github.com/terry960302/go-clean-media-upload-server/usecase"
)

type ImageMetadataController struct {
	domain.ImageMetadataContrller
	imageUsecase usecase.ImageMetadataUsecase
}

func NewImageMetadataController(imgUsecase usecase.ImageMetadataUsecase) *ImageMetadataController {
	ctrl := &ImageMetadataController{
		imageUsecase: imgUsecase,
	}
	ctrl.ImageMetadataContrller = interface{}(ctrl).(*ImageMetadataController)
	return ctrl
}

// UploadImages godoc
// @Summary upload images
// @Description upload images
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [post]

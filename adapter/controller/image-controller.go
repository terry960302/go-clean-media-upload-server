package controller

import (
	"media-upload-server/domain"
	"media-upload-server/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
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

func (i *ImageMetadataController) UploadImages(c echo.Context) error {

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	fileHeaders := form.File["files"]

	urls := i.imageUsecase.UploadImages(fileHeaders)

	return c.JSONPretty(http.StatusOK, map[string]interface {
	}{
		"data":  urls,
		"error": nil,
	}, "  ")
}

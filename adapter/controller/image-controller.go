package controller

import (
	"media-upload-server/usecase"

	"github.com/labstack/echo/v4"
)

type ImageMetadataController struct {
	imageUsecase usecase.ImageMetadataUsecase
}

func NewImageMetadataController(imgUsecase usecase.ImageMetadataUsecase) *ImageMetadataController {
	return &ImageMetadataController{
		imageUsecase: imgUsecase,
	}
}

func (i *ImageMetadataController) UploadImages(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	fileHeaders := form.File["files"]

	urls, errs := i.imageUsecase.UploadImages(fileHeaders)
}

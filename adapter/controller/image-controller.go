package controller

import (
	"net/http"

	"github.com/terry960302/go-clean-media-upload-server/domain"
	"github.com/terry960302/go-clean-media-upload-server/usecase"

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

// UploadImages godoc
// @Summary upload images
// @Description upload images
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [post]
func (i *ImageMetadataController) UploadImages(c echo.Context) error {

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	fileHeaders := form.File["files"]

	urls := i.imageUsecase.UploadImages(fileHeaders, c)

	return c.JSONPretty(http.StatusOK, map[string]interface {
	}{
		"data":  urls,
		"error": nil,
	}, "  ")
}

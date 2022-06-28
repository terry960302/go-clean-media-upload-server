package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

func (m *MediaMetadataController) UploadFiles(c echo.Context) error {

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	fileHeaders := form.File["files"]

	urls := m.mediaUsecase.UploadFiles(fileHeaders, c)

	return c.JSONPretty(http.StatusOK, map[string]interface {
	}{
		"data":  urls,
		"error": nil,
	}, "  ")
}

package domain

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MediaMetadata struct {
	gorm.Model
	MediaType string `json:"media_type"` // image, video
	Url       string `json:"url"`
}

// media > image, video

type MediaMetadataRepository interface {
	GetAll() ([]MediaMetadata, error)
	Create(media MediaMetadata) (uint, error)
}
type MediaMetadataUsecase interface {
	GetAllMedia() ([]MediaMetadata, error)
}

type MediaMetadataController interface {
	GetAllMedia(c echo.Context) error
}

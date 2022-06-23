package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type MediaMetadata struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Url       string     `json:"url"`
	CreatedAt *time.Time `json:"created_at"`
}

// media > image, video

type MediaMetadataRepository interface {
	GetAll() ([]MediaMetadata, error)
}
type MediaMetadataUsecase interface {
	GetAllMedia() ([]MediaMetadata, error)
}

type MediaMetadataController interface {
	GetAllMedia(c echo.Context) error
}

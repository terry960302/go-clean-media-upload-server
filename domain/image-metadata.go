package domain

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ImageMetadata struct {
	gorm.Model
	MediaID uint          `json:"mediaId"`
	Media   MediaMetadata `gorm:"foreignKey:MediaID"`
	Width   string        `json:"width"`
	Height  string        `json:"height"`
	Format  string        `json:"format"`
	Volume  string        `json:"volume"`
}

type ImageMetadataRepository interface {
	GetAll() ([]ImageMetadata, error)
	Create(image ImageMetadata) (uint, error)
}

type ImageMetadataUsecase interface {
	UploadImages(fileHeaders []*multipart.FileHeader, ctx echo.Context) []string
}

type ImageMetadataContrller interface {
	UploadImages(c echo.Context) error
}

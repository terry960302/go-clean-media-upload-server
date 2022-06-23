package domain

import (
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type ImageMetadata struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	MediaId   uint       `json:"mediaId"`
	Width     string     `json:"width"`
	Height    string     `json:"height"`
	Format    string     `json:"format"`
	Volume    string     `json:"volume"`
	CreatedAt *time.Time `json:"created_at"`
}

type ImageMetadataRepository interface {
	GetAll() ([]ImageMetadata, error)
}

type ImageMetadataUsecase interface {
	UploadImages(fileHeaders []*multipart.FileHeader, ctx echo.Context) []string
}

type ImageMetadataContrller interface {
	UploadImages(c echo.Context) error
}

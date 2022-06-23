package main

import (
	"media-upload-server/adapter/repository"
	"media-upload-server/config"
	"media-upload-server/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	imageRepo := repository.NewImageMetadataRepository(&gorm.DB{})
	imageUsecase := usecase.NewImageUsecase(*imageRepo)
	e.POST("/upload", imageUsecase.UploadImages)

	e.Logger.Fatal(e.Start(":" + config.C.Server.Port))
}

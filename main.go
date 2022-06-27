package main

import (
	"github.com/terry960302/go-clean-media-upload-server/adapter/controller"
	"github.com/terry960302/go-clean-media-upload-server/adapter/repository"
	"github.com/terry960302/go-clean-media-upload-server/config"
	"github.com/terry960302/go-clean-media-upload-server/domain"
	infrastructure "github.com/terry960302/go-clean-media-upload-server/infrastructure/datastore"
	"github.com/terry960302/go-clean-media-upload-server/infrastructure/router"
	"github.com/terry960302/go-clean-media-upload-server/usecase"

	"github.com/labstack/echo/v4"
)

// @title Go-Clean-Media-Upload-Server
// @version 1.0
// @description Clean Arch + Media Upload
// @host localhost:8080
// @BasePath /
func main() {
	config.ReadConfig()
	db := infrastructure.NewPostgresql()
	db.AutoMigrate(&domain.MediaMetadata{}, &domain.ImageMetadata{}, &domain.VideoMetadata{})

	appRepo := repository.NewAppRepository(db)
	appUsecase := usecase.NewAppUsecase(appRepo.ImageRepo, appRepo.MediaRepo)
	appCtrl := controller.NewAppController(appUsecase.ImageUsecase, appUsecase.MediaUsecase)

	e := echo.New()
	e = router.NewRouter(e, *appCtrl)
	e.Logger.Fatal(e.Start(":" + config.C.Server.Port))
}

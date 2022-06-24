package main

import (
	"media-upload-server/adapter/controller"
	"media-upload-server/adapter/repository"
	"media-upload-server/config"
	infrastructure "media-upload-server/infrastructure/datastore"
	"media-upload-server/infrastructure/router"
	"media-upload-server/usecase"

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
	appRepo := repository.NewAppRepository(db)
	appUsecase := usecase.NewAppUsecase(appRepo.ImageRepo, appRepo.MediaRepo)
	appCtrl := controller.NewAppController(appUsecase.ImageUsecase, appUsecase.MediaUsecase)

	e := echo.New()
	e = router.NewRouter(e, *appCtrl)
	e.Logger.Fatal(e.Start(":" + config.C.Server.Port))

}

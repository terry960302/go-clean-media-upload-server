package main

import (
	"media-upload-server/config"
	infrastructure "media-upload-server/infrastructure/datastore"

	"github.com/labstack/echo/v4"
)

func main() {

	db := infrastructure.NewPostgresql()

	e := echo.New()
	// e = router.NewRouter(e, r.NewAppController())

	e.Logger.Fatal(e.Start(":" + config.C.Server.Port))
}

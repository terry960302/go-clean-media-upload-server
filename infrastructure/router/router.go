package router

import (
	"github.com/terry960302/go-clean-media-upload-server/adapter/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/media/upload", c.MediaCtrl.UploadFiles)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e
}

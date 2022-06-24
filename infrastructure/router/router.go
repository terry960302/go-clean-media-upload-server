package router

import (
	"media-upload-server/adapter/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/images/upload", c.ImageCtrl.UploadImages)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e
}

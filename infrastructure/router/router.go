package router

import (
	"media-upload-server/adapter/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/media", func(context echo.Context) error { return c.Media.GetUsers(context) })
	// e.POST("/users", func(context echo.Context) error { return c.User.CreateUser(context) })

	return e
}

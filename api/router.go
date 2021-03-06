package api

import (
	contentV1 "clean-hexa/api/v1/content"
	contentV2 "clean-hexa/api/v2/content"

	auth "clean-hexa/api/v1/auth"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	ContentV1Controller *contentV1.Controller
	ContentV2Controller *contentV2.Controller

	AuthController *auth.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	contentV1 := e.Group("/v1/content")
	// contentV1.Use(middleware.JWTMiddleware())
	contentV1.GET("", controller.ContentV1Controller.GetAll)
	contentV1.POST("", controller.ContentV1Controller.Create)

	contentV2 := e.Group("/v2/content")
	contentV2.GET("", controller.ContentV2Controller.GetAll)
	contentV2.POST("", controller.ContentV2Controller.Create)

	auth := e.Group("/v1/auth")
	auth.POST("", controller.AuthController.Auth)
}

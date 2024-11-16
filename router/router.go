package router

import (
	"device-manager/controller"
	"device-manager/middleware"
	"device-manager/router/api"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo            *echo.Echo
	RobotController controller.RobotController
}

func (a *API) NewRouter() {

	apiGroup := a.Echo.Group("/api/v1")
	apiGroup.Use(middleware.APIKeyAuthMiddleware)
	// group
	api.RegisterRobotRoutes(apiGroup, a.RobotController)

}

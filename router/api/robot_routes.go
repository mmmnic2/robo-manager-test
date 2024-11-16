package api

import (
	"device-manager/controller"
	"github.com/labstack/echo/v4"
)

func RegisterRobotRoutes(group *echo.Group, robotController controller.RobotController) {
	robots := group.Group("/robot")
	robots.GET("/robots", robotController.GetRobots)
}

package main

import (
	"device-manager/config"
	"device-manager/controller"
	"device-manager/repository"
	"device-manager/router"
	"device-manager/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB                   = config.ConnectDatabase()
	robotRepository repository.RobotRepository = repository.NewRobotRepository(db)
	robotService    service.RobotService       = service.NewRobotService(robotRepository)
	robotController controller.RobotController = controller.NewRobotController(robotService)
)

func main() {
	defer config.CloseDatabase(db)
	e := echo.New()
	var api = router.API{Echo: e, RobotController: robotController}
	api.NewRouter()
	e.Logger.Fatal(e.Start(":1323"))

}

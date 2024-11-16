package controller

import (
	"device-manager/dto"
	"device-manager/service"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type RobotController interface {
	GetRobots(c echo.Context) error
}

type robotController struct {
	robotService service.RobotService
}

func (r *robotController) GetRobots(c echo.Context) error {
	req := dto.RobotRequest{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, dto.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request",
			Data:       nil,
		})
	}
	res, err := r.robotService.GetRobots(req)
	if err != nil {
		var e *dto.ErrorResponse
		if errors.As(err, &e) {
			return c.JSON(e.StatusCode, e)
		}
		return c.JSON(http.StatusInternalServerError, &dto.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error",
		})
	}
	return c.JSON(http.StatusOK, dto.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       res,
	})
}

func NewRobotController(robotService service.RobotService) RobotController {
	return &robotController{robotService: robotService}
}

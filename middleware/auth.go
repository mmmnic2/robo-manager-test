package middleware

import (
	"device-manager/dto"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const APIKEY = "UZXJfSslakmViv64UysvxKWPcYBjVPUN"

func APIKeyAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")

		const prefix = "ApiKey "
		if !strings.HasPrefix(authHeader, prefix) {
			return c.JSON(http.StatusUnauthorized, dto.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized: Invalid Authorization format",
			})
		}

		apiKey := strings.TrimPrefix(authHeader, prefix)

		if apiKey != APIKEY {
			return c.JSON(http.StatusForbidden, dto.Response{
				StatusCode: http.StatusForbidden,
				Message:    "Forbidden: Invalid API Key",
			})
		}

		return next(c)
	}
}

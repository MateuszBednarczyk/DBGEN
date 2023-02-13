package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Ok")
}

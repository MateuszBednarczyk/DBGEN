package handlers

import (
	"DBGEN/src/pkg/entities"
	"DBGEN/src/pkg/services"
	"github.com/labstack/echo"
	"net/http"
)

func GenerateScheme(c echo.Context) error {
	var scheme entities.Scheme
	err := c.Bind(&scheme)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}
	return c.String(http.StatusOK, services.GenerateScheme(&scheme))
}

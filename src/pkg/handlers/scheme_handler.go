package handlers

import (
	"DBGEN/src/pkg/entities"
	"DBGEN/src/pkg/services"
	"github.com/labstack/echo"
	"net/http"
)

func GenerateScheme(c echo.Context) error {
	var database entities.Database
	err := c.Bind(&database)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad request")
	}
	return c.JSON(http.StatusOK, services.GenerateDatabase(&database))
}

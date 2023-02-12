package main

import (
	"DBGEN/src/pkg/handlers"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.GET("/health", handlers.HealthCheck)

	e.Logger.Fatal(e.Start("localhost:8000"))
}

package main

import (
	"github.com/labstack/echo"

	"DBGEN/src/pkg/handlers"
)

const (
	apiVersion = "api/v1/"
)

func main() {

	e := echo.New()
	e.GET(apiVersion+"health", handlers.HealthCheck)
	e.POST(apiVersion+"creator/table", handlers.GenerateScheme)

	e.Logger.Fatal(e.Start("localhost:8000"))
}

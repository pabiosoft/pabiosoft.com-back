package main

import (
	"github.com/labstack/echo/v4"
	"pabiosoft/routes"
)

func main() {
	e := echo.New()

	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
 
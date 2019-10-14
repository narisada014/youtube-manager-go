package main

import (
	"github.com/labstack/echo"
	"youtube-manager-go/routes"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	routes.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}
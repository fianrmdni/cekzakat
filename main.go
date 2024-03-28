package main

import (
	"github.com/labstack/echo/v4"
	"cekzakat/routes"
)

func main() {
	e := echo.New()

	routes.RegisterRoutes(e)
	e.Logger.Fatal(e.Start("localhost:1234"))

}

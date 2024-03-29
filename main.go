package main

import (
	"cekzakat/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.RegisterRoutes(e)
	e.Logger.Fatal(e.Start("localhost:1234"))

}

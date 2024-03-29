package routes

import (
	"github.com/labstack/echo/v4"
	appHandler "cekzakat/handlers"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/register", appHandler.RegisterHandler)
    e.POST("/hitung-zakat", appHandler.ZakatHandler)
}
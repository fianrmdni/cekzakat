package routes

import (
	"github.com/labstack/echo/v4"
	appHandler "cekzakat/handlers"
)

func RegisterRoutes(e *echo.Echo) {
    e.POST("/hitung-zakat", appHandler.ZakatHandler)
}
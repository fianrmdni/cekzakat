package routes

import (
	appHandler "cekzakat/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	mainApp := e.Group("/cekzakat")
	mainApp.POST("/hitung-zakat", appHandler.ZakatHandler)

}

package routes

import (
	appHandler "cekzakat/handlers"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	mainApp := e.Group("/cekzakat")

	mainApp.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost", "https://cekzakat.vercel.app", "http://127.0.0.1"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	
	mainApp.POST("/hitung-zakat", appHandler.ZakatHandler)
	mainApp.POST("/register", appHandler.RegisterHandler)
	mainApp.POST("/login", appHandler.LoginHandler)
	mainApp.POST("/logout", appHandler.LogoutHandler)

}

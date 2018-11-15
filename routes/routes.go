package routes

import (
	"twilioXML/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//Run func
func Run() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	// Routes
	e.POST("/twixml", controller.TwiXML)
	// Start server
	e.Logger.Fatal(e.Start(":8500"))
}

package routes

import (
	"twilio/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//Run func
func Run() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.GET("/getxml", controller.GetXML)
	// Start server
	e.Logger.Fatal(e.Start(":8500"))
}

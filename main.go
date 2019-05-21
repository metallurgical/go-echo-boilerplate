package main

import (
	"github.com/labstack/gommon/log"
	"github.com/metallurgical/go-echo-boilerplate/config"
	"github.com/metallurgical/go-echo-boilerplate/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/joho/godotenv"
)

func main() {
	e := echo.New()

	// Load environment file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load config files
	config.AppNew()

	// Global middleware for all routes
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS handlers
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Register both web and api routes.
	routes.DefineApiRoute(e)
	routes.DefineWebRoute(e)

	// Start server to listen to port 1200
	e.Logger.Fatal(e.Start(":1200"))
}
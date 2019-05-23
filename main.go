package main

import (
	_ "fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/metallurgical/go-echo-boilerplate/database"
	"github.com/metallurgical/go-echo-boilerplate/routes"
	_ "net/http"
)

func main() {
	// Load environment file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databasePool := database.ConnectMYSQL().(*database.DatabaseProviderConnection)

	// Define API wrapper
	api := echo.New()
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())
	// CORS middleware for API endpoint.
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	routes.DefineApiRoute(api, databasePool)

	// Define WEB wrapper
	web := echo.New()
	web.Use(middleware.Logger())
	web.Use(middleware.Recover())
	routes.DefineWebRoute(web)

	server := echo.New()
	server.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()

		if req.URL.Path[:4] == "/api" {
			api.ServeHTTP(res, req)
		} else {
			web.ServeHTTP(res, req)
		}
		return
	})

	// Start server to listen to port 1200
	server.Logger.Fatal(server.Start(":1200"))
}

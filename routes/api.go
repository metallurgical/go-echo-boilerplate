package routes

import (
	_ "fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	AuthController "github.com/metallurgical/go-echo-boilerplate/controllers/api/auth"
	UserController "github.com/metallurgical/go-echo-boilerplate/controllers/api/users"
	"github.com/metallurgical/go-echo-boilerplate/database"
)

type (
	Routes struct {
		Auth  string
		Users string
	}
)

func registerRoutes() Routes {
	return Routes{
		Auth:  "/auth",
		Users: "/users",
	}
}

func DefineApiRoute(e *echo.Echo, connection database.DatabaseProvider) {
	// Group base Api wrapper into api/v1 prefix.
	api := e.Group("/api")
	routes := registerRoutes()

	func() {
		// Wrap v1 api into its own isolated section.
		v1 := api.Group("/v1")
		// Login, register, reset password, forgot password routes.
		// /auth/*
		auth := v1.Group(routes.Auth)
		auth.POST("/login", AuthController.Login(connection)).Name = "auth.login"
		auth.POST("/password/reset", AuthController.ResetPassword(connection)).Name = "auth.password.reset"

		// Authenticated user only can access these routes.
		v1.Use(authenticated)
		v1.Use(middleware.JWT([]byte("secret")))
		// /users/*
		user := v1.Group(routes.Users)
		user.GET("", UserController.Index).Name = "users.index"
	}()
}

// Check for Authorization header exist or not
// throw Unauthorized if not supply
func authenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if _, ok := c.Request().Header["Authorization"]; !ok {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

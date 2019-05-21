package routes

import (
	"github.com/labstack/echo"
	UserController "github.com/metallurgical/go-echo-boilerplate/controllers/api/users"
)

func DefineApiRoute(e *echo.Echo) {
	// Group base Api wrapper into api/v1 prefix.
	api := e.Group("/api")

	func() {
		// Wrap v1 api into its own isolated section.
		v1 := api.Group("/v1")
		user := v1.Group("/users")
		user.GET("", UserController.GetAll).Name = "users.index"
	}()
}
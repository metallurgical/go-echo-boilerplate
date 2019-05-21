package routes

import (
	"github.com/labstack/echo"
	"net/http"
)

func DefineWebRoute(e *echo.Echo) {
	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello my friend from web")
	})
}
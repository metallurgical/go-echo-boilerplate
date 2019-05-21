package users

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetAll(c echo.Context) error {
	return c.String(http.StatusOK, "All users from api")
}
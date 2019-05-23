package auth

import (
	"github.com/labstack/echo"
	"github.com/metallurgical/go-echo-boilerplate/database"
	"github.com/metallurgical/go-echo-boilerplate/mail"
	"github.com/metallurgical/go-echo-boilerplate/models"
	"net/http"
)

// Reset existing password and send email to the recipient.
func ResetPassword(connection database.DatabaseProvider) func(ctx echo.Context) error {
	return func(c echo.Context) error {
		email := c.FormValue("email")
		db = connection.(*database.DatabaseProviderConnection).Db

		if ok := (&models.User{Db: db}).IsUserExistByEmail(email); !ok {
			return echo.NewHTTPError(403, "No user matched with email address.")
		}

		if err := mail.SendResetPasswordMail(email); err != nil {
			return echo.NewHTTPError(500, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]string{
			"status":  "Success",
			"message": "Successful sent email",
		})
	}
}

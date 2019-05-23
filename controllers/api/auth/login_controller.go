package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/metallurgical/go-echo-boilerplate/database"
	"github.com/metallurgical/go-echo-boilerplate/models"
	"net/http"
	"time"
)

// Get login information as well as user JWT token.
func Login(connection database.DatabaseProvider) func(ctx echo.Context) error {
	return func(c echo.Context) error {
		var (
			email, password string
		)

		email = c.FormValue("email")
		password = c.FormValue("password")

		// Make a checking in database instead
		if email == "" || password == "" {
			return echo.NewHTTPError(403, "Please provide email and password credentials")
		}

		db = connection.(*database.DatabaseProviderConnection).Db
		if ok := (&models.User{Db: db}).IsUserExistByEmailPassword(email, password); !ok {
			return echo.ErrUnauthorized
		}
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)
		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "emi"
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
}

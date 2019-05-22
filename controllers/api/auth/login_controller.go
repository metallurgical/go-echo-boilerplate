package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/metallurgical/go-echo-boilerplate/database"
	"github.com/metallurgical/go-echo-boilerplate/models"
	"net/http"
	"time"
	"golang.org/x/crypto/bcrypt"
)

var (
	db    *gorm.DB
	count int
	user  models.User
)

func Login(connection database.DatabaseProvider) func(ctx echo.Context) error {
	return func(c echo.Context) error {
		var (
			email, password string
		)

		email = c.FormValue("email")
		password = c.FormValue("password")

		// Make a checking in database instead
		if email == "" || password == "" {
			return echo.ErrUnauthorized
		}

		db = connection.(*database.DatabaseProviderConnection).Db
		if ok := isUserExist(email, password); !ok {
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

func isUserExist(email, password string) bool {
	db.Where("email = ?", email).First(&user).Count(&count)
	if count == 0 {
		return false
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		return false;
	}
	return true
}

package service

import (
	"os"
	"time"

	"github.com/anukoon4/e-commerce-service/models"
	"github.com/dgrijalva/jwt-go"
)

func GenToken(req *models.UserModel) string {
	durationAccessToken, _ := time.ParseDuration(os.Getenv("TTL_ACCESSTOKEN"))
	tokenExpiredAt := time.Now().Add(durationAccessToken)
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = req.Email
	claims["user_id"] = req.ID
	claims["phone"] = req.Phone
	claims["name"] = req.Name
	claims["exp"] = tokenExpiredAt.Unix() // Token expiry

	// Generate the token
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return tokenString

}

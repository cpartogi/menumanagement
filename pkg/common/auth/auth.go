package auth

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken create JWT authentication token for specific user
func GenerateToken(userID int64, privateKey string, expireDuration string) (string, error) {
	var tokenString string
	duration, err := time.ParseDuration(expireDuration)
	if err != nil {
		log.Println("[JWT] invalid expire format:", err)
		return tokenString, err
	}
	tNow := time.Now()
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"exp": tNow.Add(duration).Unix(),
		"iat": tNow.Unix(),
		"sub": userID,
	}
	tokenString, err = token.SignedString([]byte(privateKey))
	return tokenString, err
}

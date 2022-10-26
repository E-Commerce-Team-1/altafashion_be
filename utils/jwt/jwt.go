package jwt

import (
	"altafashion_be/config"
	"fmt"

	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var key string

func InitJWT(c *config.AppConfig) {
	key = c.JWSecret
}

func GenerateToken(id uint) string {
	claim := make(jwt.MapClaims)
	claim["authorized"] = true
	claim["id"] = id
	claim["expired"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	str, err := token.SignedString([]byte(key))
	if err != nil {
		return ""
	}
	return str
}

func ExtractTokenProd(c echo.Context) uint {
	token := c.Get("user").(*jwt.Token)
	if token.Valid {
		claim := token.Claims.(jwt.MapClaims)
		fmt.Print(uint(claim["id"].(float64)))
		return uint(claim["id"].(float64))
	}
	return 0
}

func GenerateJWTToken(id uint) (string, error) {

	claims := make(jwt.MapClaims)
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	str, err := token.SignedString([]byte(key))

	if err != nil {
		return "", err
	}

	return str, nil
}

func ExtractToken(c echo.Context) (uint, int64) {
	token := c.Get("user").(*jwt.Token)
	// log.Println("\n\n\nisi token\n", token, "\n\n")
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		return uint(claims["id"].(float64)), int64(claims["exp"].(float64))
	}

	return 0, 0
}

func ExtractIdToken(c echo.Context) uint {
	token := c.Get("user").(*jwt.Token)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		return uint(claims["id"].(float64))
	}

	return 0
}

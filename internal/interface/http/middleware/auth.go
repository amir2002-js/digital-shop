package middleware

import (
	"errors"
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/jwtToken"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
	"time"
)

var jwtSecret []byte

func init() {
	key := os.Getenv("ACCESS_TOKEN")
	if key == "" {
		panic("ACCESS_TOKEN not found in env")
	}
	jwtSecret = []byte(key)
}

func Auth(c *fiber.Ctx) error {
	// گرفتن توکن از هدر
	auth := c.Get("Authorization")

	// گرفتن توکن اصلی
	if !strings.HasPrefix(auth, "Bearer ") {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "invalid format"})
	}
	tokenSTR := strings.TrimSpace(auth[7:])
	if tokenSTR == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "no found token"})
	}

	//ساختار توکن
	var claim jwtToken.AccessTokenClaims

	token, err := jwt.ParseWithClaims(tokenSTR, &claim, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "invalid or expired token"})
	}

	if claim.ExpiresAt == nil || claim.ExpiresAt.Before(time.Now()) {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "token expired"})
	}

	if !token.Valid {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "invalid token"})
	}

	c.Locals("user_id", claim.UserID)
	c.Locals("role", claim.Role)
	return c.Next()
}

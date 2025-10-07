package middleware

import (
	"fmt"
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/jwtToken"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
	"time"
)

func Auth(c *fiber.Ctx) error {
	// گرفتن توکن از هدر
	auth := c.Get("Authorization")

	// گرفتن توکن اصلی
	tokenSTR := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
	if tokenSTR == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "no found token"})
	}

	//ساختار توکن
	var claim jwtToken.AccessTokenClaims

	secretAccessTkn := os.Getenv("ACCESS_TOKEN")
	if secretAccessTkn == "" {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "can't find access token key"})
	}
	token, err := jwt.ParseWithClaims(tokenSTR, &claim, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token is invalid")
		}
		return []byte(secretAccessTkn), nil
	})
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	if claim.ExpiresAt != nil && claim.ExpiresAt.Before(time.Now()) {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "token expired"})
	}

	if !token.Valid {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "invalid token"})
	}

	c.Locals("user_id", claim.UserID)
	c.Locals("role", claim.Role)
	return c.Next()
}

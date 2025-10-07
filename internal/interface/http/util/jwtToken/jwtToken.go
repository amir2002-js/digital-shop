package jwtToken

import (
	"errors"
	"github.com/amir2002-js/digital-shop/internal/domain/users"
	"github.com/form3tech-oss/jwt-go"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type AccessTokenClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt2.RegisteredClaims
}

func GenerateJWTRefreshTkn(user *users.User) (strToken string, err error) {
	secretRefreshTkn := os.Getenv("REFRESH_TOKEN")
	if secretRefreshTkn == "" {
		return "", errors.New("REFRESH_TOKEN environment variable not set")
	}

	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     now.UTC().Add(time.Hour * 24 * 15).Unix(),
		"iat":     now.UTC().Unix(),
		"iss":     "digital-shop",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err = token.SignedString([]byte(secretRefreshTkn))
	return
}

func GenerateJWTAccessTkn(user *users.User) (strToken string, err error) {
	secretAccessTkn := os.Getenv("ACCESS_TOKEN")
	if secretAccessTkn == "" {
		return "", errors.New("ACCESS_TOKEN environment variable not set")
	}

	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     now.Add(time.Minute * 5).Unix(),
		"iat":     now.Unix(),
		"iss":     "digital-shop",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err = token.SignedString([]byte(secretAccessTkn))
	return
}

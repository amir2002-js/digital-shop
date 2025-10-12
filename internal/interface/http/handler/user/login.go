package usersHandler

import (
	"context"
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/jwtToken"
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/password"
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/returnsHandler"
	"github.com/gofiber/fiber/v2"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
	"strings"
	"time"
)

func (handler *UsersHandler) Login(c *fiber.Ctx) error {
	// اطلاعات کاربر (ورودی)
	entryUser := struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}{}

	// گرفتن اطلاعات از api
	if err := c.BodyParser(&entryUser); err != nil {
		return returnsHandler.CanNotBinding(c)
	}

	// چک کردن اطلاعات و تمیز کردن اطلاعات (xss) و حذف فاصله
	entryUser.Email = strings.TrimSpace(bluemonday.StrictPolicy().Sanitize(entryUser.Email))

	// اعتبار سنجی
	err := handler.validate.Struct(&entryUser)
	if err != nil {
		return returnsHandler.InvalidationData(c, err)
	}

	// پیدا کردن user
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	foundedUser, err := handler.h.IsEmailExist(ctx, entryUser.Email)
	if err != nil {
		return returnsHandler.InternalError(c, err)
	}

	if foundedUser == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "not found email or not match password", "message": "email address not found or password not match"})
	}

	// ک کردن پسورد
	isMatch, err := password.VerifyPassword(entryUser.Password, foundedUser.HashedPass)
	if err != nil {
		return returnsHandler.InternalError(c, err)
	}

	if !isMatch {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "not found email or not match password", "message": "email address not found or password not match"})
	}

	// ساخت توکن accessTkn
	accessTkn, err := jwtToken.GenerateJWTAccessTkn(foundedUser)
	if err != nil {
		return returnsHandler.InternalError(c, err)
	}

	// ساخت توکن refreshTkn
	refreshTkn, err := jwtToken.GenerateJWTRefreshTkn(foundedUser)
	if err != nil {
		return returnsHandler.InternalError(c, err)
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"data": foundedUser, "access_token": accessTkn, "refresh_token": refreshTkn})
}

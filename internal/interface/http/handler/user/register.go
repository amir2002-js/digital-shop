package usersHandler

import (
	"context"
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/jwtToken"
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/password"
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/whoIs"
	"net/http"
	"strings"
	"time"

	"github.com/amir2002-js/digital-shop/internal/domain/users"
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/returnsHandler"
	"github.com/gofiber/fiber/v2"
	"github.com/microcosm-cc/bluemonday"
)

func (handler *UsersHandler) Register(c *fiber.Ctx) error {
	// اطلاعات کاربر (ورودی)
	entryUser := struct {
		Username        string `json:"username" validate:"required,min=3,max=32"`
		Email           string `json:"email" validate:"required,email"`
		Password        string `json:"password" validate:"required,min=8"`
		ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	}{}

	// گرفتن اطلاعات از api
	if err := c.BodyParser(&entryUser); err != nil {
		return returnsHandler.CanNotBinding(c)
	}

	// چک کردن اطلاعات و تمیز کردن اطلاعات (xss)
	clearEmail := bluemonday.StrictPolicy().Sanitize(entryUser.Email)
	clearUsername := bluemonday.StrictPolicy().Sanitize(entryUser.Username)

	// حذف فاصله
	entryUser.Email = strings.TrimSpace(clearEmail)
	entryUser.Username = strings.TrimSpace(clearUsername)

	// اعتبار سنجی
	err := handler.validate.Struct(&entryUser)
	if err != nil {
		return returnsHandler.InvalidationData(c, err)
	}

	// ایا ایمیل وجود داره؟
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	userFounded, err := handler.h.IsEmailExist(ctx, entryUser.Email)
	if err != nil {
		return returnsHandler.InternalError(c, err)
	}

	if userFounded != nil {
		return returnsHandler.AlreadyExisted(c)
	}

	user := &users.User{}
	// چک کردن ادمین
	ok, err := whoIs.IsAdmin(entryUser.Password, entryUser.Email, entryUser.Username)
	if err != nil {
		if err.Error() == "credentials not set in environment variables" {
			return returnsHandler.InternalError(c, err)
		}
		return returnsHandler.AlreadyExisted(c)
	}

	// هش کردن پسورد
	hashedPass, err := password.HashPassword(entryUser.Password)
	if err != nil {
		return returnsHandler.InternalError(c, err)
	}

	// مقدار دهی یوزر جدید
	if ok {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}
	user.Username = entryUser.Username
	user.Email = entryUser.Email
	user.HashedPass = hashedPass

	// ساخت یوزر جدید
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second) // تعریف دوباره
	defer cancel()
	insertUser, err := handler.h.Register(ctx, user)
	if err != nil {
		return returnsHandler.CanNotConnectToDB(c, err)
	}

	// ساخت توکن accessTkn
	type tokenResult struct {
		Token string
		Err   error
	}
	accessCh := make(chan tokenResult, 1)
	go func() {
		accessTkn, err := jwtToken.GenerateJWTAccessTkn(insertUser)
		accessCh <- tokenResult{Token: accessTkn, Err: err}
	}()

	// ساخت توکن refreshTkn
	refreshCh := make(chan tokenResult, 1)
	go func() {
		refreshTkn, err := jwtToken.GenerateJWTRefreshTkn(insertUser)
		refreshCh <- tokenResult{Token: refreshTkn, Err: err}
	}()

	accessTknStruct := <-accessCh
	refreshTknStruct := <-refreshCh

	if accessTknStruct.Err != nil {
		return returnsHandler.InternalError(c, accessTknStruct.Err)
	}

	if refreshTknStruct.Err != nil {
		return returnsHandler.InternalError(c, refreshTknStruct.Err)
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"data": insertUser, "access_token": accessTknStruct.Token, "refresh_token": refreshTknStruct.Token})
}

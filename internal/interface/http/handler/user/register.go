package usersHandler

import (
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/password"
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/whoIs"
	"net/http"
	"strings"

	"github.com/amir2002-js/digital-shop/internal/domain/users"
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/returnsHandler"
	"github.com/gofiber/fiber/v2"
	"github.com/microcosm-cc/bluemonday"
)

func (handler *UsersHandler) Register(c *fiber.Ctx) error {
	entryUser := struct {
		Username        string `json:"username" validate:"required,min=3,max=32"`
		Email           string `json:"email" validate:"required,email"`
		Password        string `json:"password" validate:"required,min=8"`
		ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	}{}

	if err := c.BodyParser(&entryUser); err != nil {
		return returnsHandler.CanNotBinding(c)
	}

	clearEmail := bluemonday.StrictPolicy().Sanitize(entryUser.Email)
	clearUsername := bluemonday.StrictPolicy().Sanitize(entryUser.Username)

	entryUser.Email = strings.TrimSpace(clearEmail)
	entryUser.Username = strings.TrimSpace(clearUsername)

	err := handler.validate.Struct(&entryUser)
	if err != nil {
		return returnsHandler.InvalidationData(c, err)
	}

	var user *users.User

	ok, err := whoIs.IsAdmin(entryUser.Password, entryUser.Email, entryUser.Username)
	if err != nil {
		return returnsHandler.InvalidationData(c, err)
	}

	if ok {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}
	user.Username = entryUser.Username
	user.Email = entryUser.Email
	hashedPass, err := password.HashPassword(entryUser.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	user.HashedPass = hashedPass

	ctx := c.Context()
	insertUser, err := handler.h.Register(ctx, user)
	if err != nil {
		return returnsHandler.CanNotConnectToDB(c, err)
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"data": insertUser})
}

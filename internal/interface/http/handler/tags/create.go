package tagsHandler

import (
	"net/http"
	"strings"

	"github.com/amir2002-js/digital-shop/internal/interface/http/util/returnsHandler"
	"github.com/gofiber/fiber/v2"
	"github.com/microcosm-cc/bluemonday"
)

func (handler *TagsHandler) Create(c *fiber.Ctx) error {
	// چک کردن نقش
	role := c.Locals("role")
	if role != "admin" {
		return returnsHandler.NotAdminErrorAccess(c)
	}

	// خواندن از ورودی
	entryTag := struct {
		Name string `validate:"required;min=3;max=10" json:"name"`
	}{}
	err := c.BodyParser(&entryTag)
	if err != nil {
		return returnsHandler.CanNotBinding(c)
	}

	// ولیدیت مقادیر
	entryTag.Name = strings.ToLower(bluemonday.StrictPolicy().Sanitize(entryTag.Name))
	err = handler.validate.Struct(&entryTag)
	if err != nil {
		return returnsHandler.InvalidationData(c, err)
	}

	ctx := c.UserContext()
	tag, err := handler.h.CreateTag(ctx, entryTag.Name)
	if err != nil {
		return returnsHandler.CanNotConnectToDB(c, err)
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"data": tag})
}

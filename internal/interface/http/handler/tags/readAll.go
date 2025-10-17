package tagsHandler

import (
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/returnsHandler"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (handler *TagsHandler) ReadAll(c *fiber.Ctx) error {
	ctx := c.UserContext()
	tags, err := handler.h.ReadAllTag(ctx)
	if err != nil {
		return returnsHandler.CanNotConnectToDB(c, err)
	}
	if tags == nil {
		return c.Status(http.StatusOK).JSON(fiber.Map{"data": nil, "message": "data is empty"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": tags})
}

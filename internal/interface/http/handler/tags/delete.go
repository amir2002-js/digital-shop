package tagsHandler

import (
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/returnsHandler"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (handler *TagsHandler) Delete(c *fiber.Ctx) error {
	role := c.Locals("role")
	if role != "admin" {
		return returnsHandler.NotAdminErrorAccess(c)
	}

	tagId, err := c.ParamsInt("id", -1)
	if err != nil || tagId < 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "id is empty or not valid"})
	}

	ctx := c.UserContext()
	err = handler.h.DeleteByIdTag(ctx, tagId)
	if err != nil {
		return returnsHandler.CanNotConnectToDB(c, err)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": tagId, "success": true, "message": "tag deleted"})
}

package galleryHandler

import (
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/returnsHandler"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (handler *GalleryHandler) RemoveFromGallery(c *fiber.Ctx) error {
	userRole := c.Locals("role").(string)
	if userRole != "admin" {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "you are not admin"})
	}

	imgId, err := c.ParamsInt("id", -1)
	if err != nil || imgId < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid or empty id"})
	}

	ctx := c.UserContext()
	err = handler.h.RemoveFromGallery(ctx, imgId)
	if err != nil {
		return returnsHandler.CanNotConnectToDB(c, err)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"massage": "success removed from gallery"})
}

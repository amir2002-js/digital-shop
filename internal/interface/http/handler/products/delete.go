package productsHandler

import (
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/returnsHandler"
	"github.com/gofiber/fiber/v2"
)

func (handler *ProductsHandler) Delete(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "you are not admin"})
	}

	productID, err := c.ParamsInt("id", -1)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id is not valid or empty"})
	}

	ctx := c.UserContext()
	err = handler.h.Delete(ctx, productID)
	if err != nil {
		return returnsHandler.CanNotConnectToDB(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Product deleted"})
}

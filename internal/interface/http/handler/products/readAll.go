package productsHandler

import (
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/returnsHandler"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (handler *ProductsHandler) ReadAll(c *fiber.Ctx) error  {
	ctx := c.Context()
	allProducts ,err := handler.h.ReadAll(ctx)
	if err != nil {
		return returnsHandler.InternalError(c, err)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data" : allProducts})
}

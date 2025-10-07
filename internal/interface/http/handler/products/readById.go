package productsHandler

import (
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/returnsHandler"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (handler *ProductsHandler) ReadById(c *fiber.Ctx) error  {
	productID ,err := c.ParamsInt("id" , -1)
	if err != nil {
		return returnsHandler.InternalError(c,err)
	}

	if productID < 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Product Id is not valid"})
	}

	ctx := c.Context()
	product ,err := handler.h.ReadById(ctx ,productID)
	if err != nil {
		return returnsHandler.InternalError(c, err)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data" : product})
}


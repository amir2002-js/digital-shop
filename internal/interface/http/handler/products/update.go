package productsHandler

import (
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/returnsHandler"
	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
	"net/http"
)

func (handler *ProductsHandler) Update(c *fiber.Ctx) error {
	// چک کردن رول کاربر
	userRole := c.Locals("role").(string)
	if userRole != "admin" {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "you are not admin"})
	}

	// گرفتن ایدی محصول
	pID, err := c.ParamsInt("id", -1)
	if err != nil || pID < 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "id is not valid or empty"})
	}

	// ورودی محصول
	entryValues := struct {
		Price       *decimal.Decimal `json:"price" validate:"omitempty,gt=0"`
		Name        string           `json:"name" validate:"omitempty,min=3"`
		Description string           `json:"description" validate:"omitempty,min=10"`
		Discount    *decimal.Decimal `json:"discount" validate:"omitempty,gte=0,lt=100"`
		Stock       *uint            `json:"stock" validate:"omitempty,gte=0"`
	}{}
	err = c.BodyParser(&entryValues)
	if err != nil {
		return returnsHandler.CanNotBinding(c)
	}

	// محصول قدیمی
	ctx := c.Context()
	oldProduct, err := handler.h.ReadById(ctx, pID)
	if err != nil {
		return returnsHandler.CanNotConnectToDB(c, err)
	}

	// مقدار دهی
	if entryValues.Price != nil {
		oldProduct.PriceBeforeOff = *entryValues.Price
	}
	if entryValues.Stock != nil {
		oldProduct.Stock = *entryValues.Stock
	}
	if entryValues.Name != "" {
		oldProduct.Name = entryValues.Name
	}
	if entryValues.Description != "" {
		oldProduct.Description = entryValues.Description
	}
	if entryValues.Discount != nil {
		oldProduct.Discount = *entryValues.Discount
	}
	// اپدیت محصول
	ctx = c.Context()
	err = handler.h.Update(ctx, oldProduct)
	if err != nil {
		return returnsHandler.CanNotConnectToDB(c, err)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"product": oldProduct})
}

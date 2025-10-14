package productsHandler

import (
	"context"
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/returnsHandler"
	"github.com/gofiber/fiber/v2"
	"github.com/microcosm-cc/bluemonday"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
	"strings"
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
	ctx := context.Background()
	productFind, err := handler.h.ReadById(ctx, pID)

	// اپدیت قیمت
	priceValue := c.FormValue("price", "")
	if priceValue != "" {
		price, err := decimal.NewFromString(priceValue)
		if err != nil {
			return returnsHandler.CanNotBinding(c)
		}
		productFind.PriceBeforeOff = price
	}

	// اپدیت نام
	nameValue := c.FormValue("name", "")
	if nameValue != "" {
		clearName := bluemonday.StrictPolicy().Sanitize(nameValue)
		clearName = strings.TrimSpace(clearName)
		productFind.Name = clearName
	}

	// اپدیت توضیحات
	descriptionValue := c.FormValue("description", "")
	if descriptionValue != "" {
		clearDescription := bluemonday.StrictPolicy().Sanitize(descriptionValue)
		clearDescription = strings.TrimSpace(clearDescription)
		productFind.Description = clearDescription
	}

	// اپدیت تخفیف
	discountValue := c.FormValue("discount", "")
	if discountValue != "" {
		discount, err := decimal.NewFromString(discountValue)
		if err != nil {
			return returnsHandler.CanNotBinding(c)
		}
		productFind.Discount = discount
	}

	// اپدیت تعداد محصول
	stockValue := c.FormValue("stock", "")
	if stockValue != "" {
		stockInt, err := strconv.Atoi(stockValue)
		stockUint := uint(stockInt)
		if err != nil || stockUint > 100 {
			return returnsHandler.CanNotBinding(c)
		}
		productFind.Stock = stockUint
	}

	ctx = c.UserContext()
	err = handler.h.Update(ctx, productFind)
	if err != nil {
		return returnsHandler.CanNotConnectToDB(c, err)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"product": productFind})
}

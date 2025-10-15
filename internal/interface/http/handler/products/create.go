package productsHandler

import (
	"github.com/amir2002-js/digital-shop/internal/domain/products"
	producttags "github.com/amir2002-js/digital-shop/internal/domain/productsTags"
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/returnsHandler"
	"github.com/gofiber/fiber/v2"
	"github.com/microcosm-cc/bluemonday"
	"github.com/shopspring/decimal"
	"net/http"
	"strings"
)

func (handler *ProductsHandler) Create(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	if role != "admin" {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "You are not allowed to add products"})
	}

	var entryProducts = struct {
		Name        string          `json:"name" validate:"required;min=3,max=100"`
		Price       decimal.Decimal `json:"price" validate:"required;gt=0"`
		Stock       uint            `json:"stock" validate:"required;min=0"`
		Discount    decimal.Decimal `json:"discount" validate:"required;gte=0,lt=100"`
		Description string          `json:"description" validate:"required;min=1"`
		Tags        []uint          `json:"tags" validate:"required,min=1"`
	}{}

	err := c.BodyParser(&entryProducts)
	if err != nil {
		return returnsHandler.CanNotBinding(c)
	}

	clearName := strings.TrimSpace(bluemonday.StrictPolicy().Sanitize(entryProducts.Name))
	clearDescription := strings.TrimSpace(bluemonday.StrictPolicy().Sanitize(entryProducts.Description))

	entryProducts.Name = clearName
	entryProducts.Description = clearDescription

	var newProduct products.Product

	newProduct.Name = entryProducts.Name
	newProduct.Discount = entryProducts.Discount
	newProduct.Description = entryProducts.Description
	newProduct.Stock = entryProducts.Stock
	newProduct.PriceBeforeOff = entryProducts.Price

	ctx := c.UserContext()
	err = handler.h.Create(ctx, &newProduct)
	if err != nil {
		return returnsHandler.CanNotConnectToDB(c, err)
	}

	var arrTags = make([]producttags.ProductTag, len(entryProducts.Tags))
	for tagID := range entryProducts.Tags {
		ctxTag := c.UserContext()
		if find, _ := handler.h.FindTag(ctxTag, tagID); find {
			var newTag = producttags.ProductTag{
				ProductID: int(newProduct.ID),
				TagID:     tagID,
			}
			arrTags = append(arrTags, newTag)
		}
	}

	ctxAddToTags := c.UserContext()
	err = handler.h.AddToTags(ctxAddToTags, arrTags)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"data": newProduct, "error": "can't add tags to products but your product added"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"date": newProduct})
}

package galleryHandler

import (
	"github.com/amir2002-js/digital-shop/internal/domain/gallery"
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/images"
	"github.com/amir2002-js/digital-shop/internal/interface/http/util/returnsHandler"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (handler *GalleryHandler) AddToGallery(c *fiber.Ctx) error {
	userRole := c.Locals("role").(string)
	if userRole != "admin" {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "you are not admin"})
	}

	productId, err := c.ParamsInt("productId", -1)
	if err != nil || productId < 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid or empty productId"})
	}

	ctx := c.UserContext()
	galleryArr, err := handler.h.GetImageByProductId(ctx, productId)
	if err != nil {
		return returnsHandler.CanNotConnectToDB(c, err)
	}

	if len(galleryArr) <= 5 {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "can't add to gallery when number of images product is 5"})
	}

	imgURL, err := images.CreateImage(c, map[string]bool{"jpg": true, "jpeg": true}, 1024*1024*5)
	if err != nil {
		return returnsHandler.InternalError(c, err)
	}

	var newImageOfGallery gallery.Gallery

	newImageOfGallery.Url = imgURL
	newImageOfGallery.IsMain = false
	newImageOfGallery.ProductId = uint(productId)

	err = handler.h.AddToGallery(ctx, &newImageOfGallery)
	if err != nil {
		return returnsHandler.CanNotConnectToDB(c, err)
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"data": newImageOfGallery})
}

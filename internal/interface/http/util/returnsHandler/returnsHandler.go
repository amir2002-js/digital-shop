package returnsHandler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CanNotBinding(c *fiber.Ctx) error {
	return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "cannot binding"})
}

func InvalidationData(c *fiber.Ctx, err error) error {
	return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "message": "invalid data"})
}

func CanNotConnectToDB(c *fiber.Ctx, err error) error {
	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "cannot connect to database", "message": err.Error()})
}

func AlreadyExisted(c *fiber.Ctx) error {
	return c.Status(http.StatusConflict).JSON(fiber.Map{"error": "email is already exist"})
}

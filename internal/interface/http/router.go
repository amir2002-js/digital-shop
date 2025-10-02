package http

import (
	"github.com/amir2002-js/digital-shop/internal/interface/http/handler"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App, handler *handler.Handler) {

	v1 := app.Group("/api/v1")
	{
		adminContact := v1.Group("/admin-contact")
		{
			adminContact.Get("/", nil)
			adminContact.Put("/", nil)
			adminContact.Post("/", nil)
		}

		login := v1.Group("/login")
		{
			login.Post("/", nil)
		}

		register := v1.Group("/register")
		{
			register.Post("/", nil)
		}

		teammate := v1.Group("/teammate")
		{
			teammate.Put("/:id", nil)
			teammate.Get("/", nil)
			teammate.Delete("/:id", nil)
			teammate.Post("/", nil)
		}
	}
}

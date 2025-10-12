package http

import (
	"github.com/amir2002-js/digital-shop/internal/interface/http/handler"
	"github.com/amir2002-js/digital-shop/internal/interface/http/middleware"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App, h *handler.Handler) {

	v1 := app.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.Post("/login", h.User.Login)
			user.Post("/register", h.User.Register)
		}

		product := v1.Group("/product")
		{
			product.Post("/", middleware.Auth)
			product.Put("/", middleware.Auth)
			product.Get("/", h.Product.ReadAll)
			product.Get("/:id", h.Product.ReadById)
			product.Delete("/", middleware.Auth)
		}
	}
}

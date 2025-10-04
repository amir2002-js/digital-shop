package http

import (
	"github.com/amir2002-js/digital-shop/internal/interface/http/handler"
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
	}
}

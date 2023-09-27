package router

import (
	"github.com/daniu/ecommerce/app/handler"
	"github.com/gofiber/fiber/v2"
)


func AuthRoutes(auth fiber.Router)  {
	auth.Get("/login", handler.RenderLogin)
	auth.Get("/signup", handler.RenderSignup)
}

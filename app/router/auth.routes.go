package router

import (
	"github.com/daniu/ecommerce/app/handler"
	"github.com/gofiber/fiber/v2"
)


func AuthRoutes(auth fiber.Router)  {
	auth.Get("/login", handler.RenderLogin)
	auth.Post("/login", handler.Login)
	auth.Get("/signup", handler.RenderSignup)
}

package router

import (
	"github.com/daniu/ecommerce/app/handler"
	"github.com/gofiber/fiber/v2"
)


func HomeRoutes(home fiber.Router)  {
	home.Get("/", handler.RenderHome)
}

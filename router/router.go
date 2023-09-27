package router

import (
	"github.com/daniu/ecommerce/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)


func SetupRoutes(app *fiber.App)  {
	home := app.Group("/", logger.New())
	HomeRoutes(home)

	login := app.Group("/auth", logger.New())
	AuthRoutes(login)

	// Middlware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)
}

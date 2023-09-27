package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)


func SetupRoutes(app *fiber.App)  {
	home := app.Group("/", logger.New())
	HomeRoutes(home)

	login := app.Group("/auth", logger.New())
	AuthRoutes(login)

}

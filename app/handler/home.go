package handler

import "github.com/gofiber/fiber/v2"

func RenderHome(c *fiber.Ctx) error {
	err := c.Render("index",fiber.Map{
		"Title":"Hello!",
	})
	if err != nil {
		c.SendString("Render page not working")
	}
	return err
}

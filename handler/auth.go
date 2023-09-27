package handler

import "github.com/gofiber/fiber/v2"

func RenderLogin(c *fiber.Ctx) error {
	return c.Render("auth/login",fiber.Map{
		"Title":"Login",
	})
}

func RenderSignup(c *fiber.Ctx) error {
	return c.Render("auth/signup",fiber.Map{
		"Title":"Sign Up",
	})
}
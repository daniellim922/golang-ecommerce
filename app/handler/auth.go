package handler

import (
	"github.com/gofiber/fiber/v2"
)

type loginForm struct {
	Email string
	Password string
}

func RenderLogin(c *fiber.Ctx) error {
	err := c.Render("auth/login",fiber.Map{
		"Title":"Login",
	})
	if err != nil {
		c.SendString("Error Rendering login page")
	}
	return err
}

func Login(c *fiber.Ctx) error {
	data := loginForm{
		Email: c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	return c.JSON(data)
}

func RenderSignup(c *fiber.Ctx) error {
	return c.Render("auth/signup",fiber.Map{
		"Title":"Sign Up",
	})
}
package handler

import (
	"github.com/daniu/ecommerce/app/model"
	"github.com/gofiber/fiber/v2"
)


type SignupForm struct {
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
	data := model.User{
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

func Signup(c *fiber.Ctx) error {
	data := model.User{
		Username: c.FormValue("username"),
		Email: c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	return c.JSON(data)
}

package main

import (
	"log"

	"github.com/daniu/ecommerce/config"
	"github.com/daniu/ecommerce/db"
	"github.com/daniu/ecommerce/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/template/handlebars/v2"
)


func main()  {
	engine := handlebars.New("./views", ".hbs")
	app := fiber.New(fiber.Config{
		Views: engine,
		ViewsLayout: "layouts/main",
	})
	
	db.ConnectDB()
	
	app.Use(helmet.New())
	app.Static("/","./public")
	router.SetupRoutes(app)

	port := config.Config("SERVER_PORT")
	log.Fatal(app.Listen(port))
}
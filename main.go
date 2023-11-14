package main

import (
	"wiai/controllers"
	"wiai/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	helpers.Establish()
	helpers.Migrate()
}

func main() {
	engine := html.New("./views", ".html")
	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: engine, 
	})

	app.Static("/public/", "./public")

	controllers.Index(app)
	controllers.Login(app)
	controllers.Register(app)
	controllers.User(app)

	app.Listen("127.0.0.1:3000")
}
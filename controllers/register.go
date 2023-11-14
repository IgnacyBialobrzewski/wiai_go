package controllers

import "github.com/gofiber/fiber/v2"

func register(c *fiber.Ctx) error {
	return c.Render("register", nil, "fragments/layout")
}

func Register(a *fiber.App) {
	a.Get("/register", register)
}
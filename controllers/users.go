package controllers

import (
	"errors"
	"log"
	"net/http"
	"wiai/helpers"
	"wiai/models"

	"github.com/gofiber/fiber/v2"
)

func Users(a *fiber.App) {
	a.Post("/users", registerUser)
}

func registerUser(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := models.CreateUser(helpers.Db, username, password)

	if err != nil {
		log.Printf("failed to create user: %+v", err)

		if errors.Is(err, models.ErrUserAlreadyExists) {
			return c.SendString("User already exists!")
		}

		return c.SendString("Error occurred while registering")
	}

	log.Println("Registered: " + user.Username)

	return c.Status(http.StatusCreated).SendString("Registered successfully!")
}

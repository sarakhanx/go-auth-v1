package userControllers

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/sarakhanx/go-auth-v1/config"
	"github.com/sarakhanx/go-auth-v1/models"
	"github.com/sarakhanx/go-auth-v1/queries"
)

func Helloworld(c fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func Signup(c fiber.Ctx) error {
	conn := config.InitDB()
	var data models.User
	err := c.Bind().Body(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid data in request")
	}

	_, err = conn.Exec(context.Background(), queries.SignupNewUser, data.Name, data.Username, data.Password, data.Email, data.Roles)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error while create user")
	}
	return c.JSON(data)
}

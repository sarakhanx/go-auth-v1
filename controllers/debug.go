package userControllers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/sarakhanx/go-auth-v1/config"
	"github.com/sarakhanx/go-auth-v1/models"
)

func Debug(c fiber.Ctx) error {
	conn := config.InitDB()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select * from users;")
	if err != nil {
		log.Println("Error executing query:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error executing query"})
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.Email, &user.Roles)
		if err != nil {
			log.Println("Error scanning row:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning row"})
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error iterating rows"})
	}

	return c.JSON(fiber.Map{"data": users})
}
func Helloworld(c fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

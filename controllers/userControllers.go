package userControllers

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v4"
	"github.com/sarakhanx/go-auth-v1/config"
	validate "github.com/sarakhanx/go-auth-v1/middlewares"
	"github.com/sarakhanx/go-auth-v1/models"
	"github.com/sarakhanx/go-auth-v1/queries"
	"github.com/sarakhanx/go-auth-v1/utils"
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// NOTE -  signup
func Signup(c fiber.Ctx) error {
	conn := config.InitDB()

	var data models.User

	err := c.Bind().Body(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid data in request")
	}
	//EXPLAIN - check all data is not empty
	if strings.TrimSpace(data.Name) == "" ||
		strings.TrimSpace(data.Username) == "" ||
		strings.TrimSpace(data.Password) == "" ||
		strings.TrimSpace(data.Email) == "" ||
		strings.TrimSpace(data.Roles) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "All fields are required"})
	}
	//EXPLAIN - check email valiadte
	if !validate.IsValidEmail(data.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid email"})
	}
	//check lenght of password
	if len(data.Password) < 8 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password must be at least 8 characters"})
	}

	//EXPLAIN - check if user exists
	var existingUser models.User
	err = conn.QueryRow(context.Background(), queries.CheckUserExists, data.Username, data.Email).Scan(&existingUser.Username, &existingUser.Email)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username or email already exists"})
	} else if err.Error() != "no rows in result set" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error while checking existing user"})
	}

	//EXPLAIN - encode password
	hashedPassword, err := HashedPassword(data.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error while hashing password")
	}
	//EXPLAIN - start query
	_, err = conn.Exec(context.Background(), queries.SignupNewUser, data.Name, data.Username, hashedPassword, data.Email, data.Roles)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Error while create user")
	}
	return c.JSON(data)
}

// NOTE - signin
func Signin(c fiber.Ctx) error {
	conn := config.InitDB()
	defer conn.Close(context.Background())
	var data models.SigninUser
	err := c.Bind().Body(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid data in request")
	}

	//EXPLAIN - check username and password
	var hashedPassword string
	err = conn.QueryRow(context.Background(), queries.SigninUser, data.Username).Scan(&hashedPassword)
	if err != nil {
		if err == pgx.ErrNoRows {
			return c.Status(fiber.StatusBadRequest).JSON("Not found user")
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error while checking user"})
	}
	//EXPLAIN - check password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(data.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON("Incorrect password")
	}

	token, err := utils.GenerateToken(data.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error while generating token"})
	}

	data.Password = ""

	return c.JSON(fiber.Map{"message": "Signin successful", "data": data, "token": token, "hashed": hashedPassword})
}

// NOTE - Signout
func Signout(c fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing authorization token"})
	}
	return c.JSON(fiber.Map{"message": "Signed Out!👋"})
}

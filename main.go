package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/sarakhanx/go-auth-v1/config"
	userRouter "github.com/sarakhanx/go-auth-v1/routes"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()
	// Database connection
	conn := config.InitDB()
	defer conn.Close(context.Background())
	//Debug if server is starting
	log.Println("Server is starting...")

	// Route
	userRouter.SetupUser(app)
	userRouter.PdfControllers(app)

	log.Println("Server is running on port 8080...")
	log.Fatal(app.Listen(":8080"))
}

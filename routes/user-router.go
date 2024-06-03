package userRouter

import (
	"github.com/gofiber/fiber/v3"
	userControllers "github.com/sarakhanx/go-auth-v1/controllers"
)

func SetupUser(app *fiber.App) {
	app.Get("/hifolks", userControllers.Helloworld)
	app.Post("/api/auth/signup", userControllers.Signup)
}

package userRouter

import (
	"github.com/gofiber/fiber/v3"
	userControllers "github.com/sarakhanx/go-auth-v1/controllers"
)

func SetupUser(app *fiber.App) {
	app.Get("/hifolks", userControllers.Helloworld)
	app.Get("/debuguser", userControllers.Debug)
	app.Post("/api/auth/signup", userControllers.Signup)
	app.Post("/api/auth/signin", userControllers.Signin)
	app.Post("/api/auth/signout", userControllers.Signout)
}

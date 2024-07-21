package userRouter

import (
	"github.com/gofiber/fiber/v3"
	auth_controllers "github.com/sarakhanx/go-auth-v1/controllers/auth-controllers"
	pdfControllers "github.com/sarakhanx/go-auth-v1/controllers/go-pdf-controler"
)

func SetupUser(app *fiber.App) {
	app.Get("/hifolks", auth_controllers.Helloworld)
	app.Get("/debuguser", auth_controllers.Debug)
	app.Post("/api/auth/signup", auth_controllers.Signup)
	app.Post("/api/auth/signin", auth_controllers.Signin)
	app.Post("/api/auth/signout", auth_controllers.Signout)
}

func PdfControllers(app *fiber.App) {
	app.Post("/api/create-sample-pdf", pdfControllers.PdfHandler)
}

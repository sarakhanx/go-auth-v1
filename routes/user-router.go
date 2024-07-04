package userRouter

import (
	"github.com/gofiber/fiber/v3"
	userControllers "github.com/sarakhanx/go-auth-v1/controllers"
	pdfControllers "github.com/sarakhanx/go-auth-v1/controllers/go-pdf-controler"
)

func SetupUser(app *fiber.App) {
	app.Get("/hifolks", userControllers.Helloworld)
	app.Get("/debuguser", userControllers.Debug)
	app.Post("/api/auth/signup", userControllers.Signup)
	app.Post("/api/auth/signin", userControllers.Signin)
	app.Post("/api/auth/signout", userControllers.Signout)
}

func PdfControllers(app *fiber.App) {
	app.Post("/api/create-sample-pdf", pdfControllers.PdfHandler)
}

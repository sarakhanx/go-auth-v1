package gopdfcontroler

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	fpdfConfig "github.com/sarakhanx/go-auth-v1/config/fpdf"
	"github.com/sarakhanx/go-auth-v1/models"
)

func PdfHandler(c fiber.Ctx) error {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exeDir := filepath.Dir(exePath)
	envPath := filepath.Join(exeDir, ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Println("Error loading .env file", err)
	}
	if err = godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file", err)
	}

	pdfDir := os.Getenv("PDF_DIR")

	var data models.PDFRequest

	err = c.Bind().Body(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid data in request")
	}
	directory := pdfDir
	var filename string = time.Now().Format("20060102_150405") + data.Title
	filepath := directory + filename + ".pdf"

	pdfData, err := fpdfConfig.PdfGenerator1(data, filepath)
	if err != nil {
		log.Println("Error generating PDF:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error generating PDF"})
	}
	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename=pdf.pdf")
	return c.Send(pdfData)
}

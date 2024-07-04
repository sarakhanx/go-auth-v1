package fpdfConfig

import (
	"bytes"
	"log"
	"os"

	"github.com/go-pdf/fpdf"
	pdfvalidations "github.com/sarakhanx/go-auth-v1/middlewares/pdf-validations"
	"github.com/sarakhanx/go-auth-v1/models"
)

func PdfGenerator1(data models.PDFRequest, filename string) ([]byte, error) {

	err := pdfvalidations.IsPdfValid((data))
	if err != nil {
		log.Printf("PDF validation Error: %v", err)
		return nil, err
	}
	//prepare a page
	pdf := fpdf.New("P", "mm", "A4", "")
	//add page
	pdf.AddPage()
	//Header
	pdf.SetFont("Arial", "B", 16)
	pdf.SetTextColor(88, 57, 39)
	pdf.Cell(40, 10, data.Title)
	pdf.Ln(10)
	//Content
	pdf.SetFont("Arial", "", 8)
	pdf.SetTextColor(0, 0, 0)
	pdf.MultiCell(0, 10, data.Content, "", "", false)
	pdf.Ln(2)
	//Author
	pdf.SetFont("Arial", "", 8)
	pdf.SetTextColor(0, 0, 0)
	pdf.Cell(40, 10, data.Author)

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Error while creating file: %v", err)
		return nil, err
	}

	defer file.Close()

	_, err = file.Write(buf.Bytes())
	if err != nil {
		log.Printf("Error while writing file: %v", err)
		return nil, err
	}
	log.Printf("PDF generated successfully : %s", filename)

	return buf.Bytes(), nil
}

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
	// Set font and colors for the table headers
	pdf.SetFont("Arial", "B", 16)
	pdf.SetFillColor(200, 200, 200) // Light grey background
	pdf.SetTextColor(0, 0, 0)       // Black text
	pdf.SetDrawColor(0, 0, 0)       // Black border

	// สร้าง Table headers
	headers := []string{"Title", "Content", "Author"}
	for _, header := range headers {
		pdf.CellFormat(60, 10, header, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	// Set font and colors for the table content
	pdf.SetFont("Arial", "", 12)
	pdf.SetFillColor(255, 255, 255) // White background
	pdf.SetTextColor(0, 0, 0)       // Black text
	pdf.SetDrawColor(0, 0, 0)       // Black border

	// Table content
	pdf.CellFormat(60, 10, data.Title, "1", 0, "C", true, 0, "")
	pdf.CellFormat(60, 10, data.Content, "1", 1, "C", true, 0, "")
	pdf.CellFormat(60, 10, data.Author, "1", 0, "C", true, 0, "")
	pdf.Ln(-1)

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

package pdfvalidations

import (
	"errors"

	"github.com/sarakhanx/go-auth-v1/models"
)

func IsPdfValid(data models.PDFRequest) error {
	if data.Title == "" {
		return errors.New("title is required")
	}
	if data.Content == "" {
		return errors.New("content is required")
	}
	if data.Author == "" {
		return errors.New("author is required")
	}

	return nil
}

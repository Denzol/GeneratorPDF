package certgenerator

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"path/filepath"
)

func Save(pdfgen *wkhtmltopdf.PDFGenerator) error {
	absPath, _ := filepath.Abs("../utils/storage/local/PDF/example.pdf")
	err := pdfgen.WriteFile(absPath)
	if err != nil {
		return err
	}
	return nil
}

package certgenerator

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GeneratePDF(r RequestPdf) (*wkhtmltopdf.PDFGenerator, error) {

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(r.Buf))

	switch r.Format {
	case "A3":
		pdfg.PageSize.Set(wkhtmltopdf.PageSizeA3)
	case "A4":
		pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	case "A5":
		pdfg.PageSize.Set(wkhtmltopdf.PageSizeA5)
	case "A7":
		pdfg.PageSize.Set(wkhtmltopdf.PageSizeA7)
	}

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg, nil
}

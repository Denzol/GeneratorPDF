package certgenerator

import (
	"errors"
	"fmt"
)

func Starter() error {
	newData := PdfData{
		Templ:   "sample.html",
		Format:  "A4",
		Title:   "Certificate Golang School",
		Student: "Khramtsov Denis",
		Course:  "Become a gopher",
		Mentors: "Pavel Gordiyanov, Mikita Viarbovikau, Sergey Shtripling",
		Date:    "08.09.2022",
	}

	tmpl, err := ParseTemplate(newData)
	if err != nil {
		errors.New("Template not prepared!")
	}

	pdfg, err := GeneratePDF(tmpl)
	if err != nil {
		return errors.New("PDF not generate!")
	}

	err = Save(pdfg)
	if err != nil {
		return errors.New("PDF not save!")
	}

	fmt.Println("Done!")

	return nil
}

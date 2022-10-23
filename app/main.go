package main

import (
	"GeneratorPDF/app/certgenerator"
	"log"
)

func main() {
	err := certgenerator.Starter()
	if err != nil {
		log.Fatal(err)
	}
}

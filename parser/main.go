package main

import (
	"bufio"
	"ddjj/parser/declaration"
	"ddjj/parser/extractor"
	"log"
	"os"
)

func main() {
	file := os.Args[1]
	data, err := os.Open(file)
	if err != nil {
		log.Fatal("Failed to open file")
	}

	scanner := bufio.NewScanner(data)
	d := &declaration.Declaration{
		Ano:         extractor.Year(scanner),
		Cedula:      extractor.Cedula(scanner),
		Nombre:      extractor.Name(scanner),
		Apellido:    extractor.Lastname(scanner),
		Institucion: extractor.Institution(scanner),
		Funcion:     extractor.JobTitle(scanner),
	}

	log.Println(d)
}

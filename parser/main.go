package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"ddjj/parser/declaration"
	"ddjj/parser/extract"
)

func main() {
	file := os.Args[1]
	data, err := os.Open(file)
	if err != nil {
		log.Fatal("Failed to open file")
	}

	// Basic Info.
	scanner := bufio.NewScanner(data)
	d := &declaration.Declaration{
		Ano:         extract.Year(scanner),
		Cedula:      extract.Cedula(scanner),
		Nombre:      extract.Name(scanner),
		Apellido:    extract.Lastname(scanner),
		Institucion: extract.Institution(scanner),
		Funcion:     extract.JobTitle(scanner),
	}

	// Deposits.
	data, _ = os.Open(file)
	scanner = bufio.NewScanner(data)
	d.Deposits = extract.Deposits(scanner)

	// Debtors.
	data, _ = os.Open(file)
	scanner = bufio.NewScanner(data)
	d.Debtors = extract.Debtors(scanner)

	print(d)
}

func print(d *declaration.Declaration) {
	fmt.Printf("Año: %d\nCedula: %d\nName: %s\nInstitution: %s\nJob: %s\n",
		d.Ano, d.Cedula, d.Nombre+" "+d.Apellido, d.Institucion, d.Funcion)

	fmt.Printf("\nDepósitos:\n")
	for _, deposit := range d.Deposits {
		fmt.Println(deposit)
	}

	fmt.Print("\nCuentas a cobrar:\n")
	for _, debtor := range d.Debtors {
		fmt.Println(debtor)
	}
}

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
		Date:        extract.Date(scanner),
		Cedula:      extract.Cedula(scanner),
		Nombre:      extract.Name(scanner),
		Apellido:    extract.Lastname(scanner),
		Institucion: extract.Institution(scanner),
		Funcion:     extract.JobTitle(scanner),
	}

	// Deposits.
	data, _ = os.Open(file)
	defer data.Close()
	scanner = bufio.NewScanner(data)
	d.Deposits = extract.Deposits(scanner)

	// Debtors.
	data, _ = os.Open(file)
	defer data.Close()
	scanner = bufio.NewScanner(data)
	d.Debtors = extract.Debtors(scanner)

	// Real state.
	data, _ = os.Open(file)
	defer data.Close()
	scanner = bufio.NewScanner(data)
	d.RealStates = extract.RealStates(scanner)

	// Vehicles
	data, _ = os.Open(file)
	defer data.Close()
	scanner = bufio.NewScanner(data)
	d.Vehicles = extract.Vehicles(scanner)

	// Agricultural activity
	data, _ = os.Open(file)
	defer data.Close()
	scanner = bufio.NewScanner(data)
	d.Agricultural = extract.Agricultural(scanner)

	// Furniture
	data, _ = os.Open(file)
	defer data.Close()
	scanner = bufio.NewScanner(data)
	d.Furniture = extract.Furniture(scanner)

	// Other assets
	data, _ = os.Open(file)
	defer data.Close()
	scanner = bufio.NewScanner(data)
	d.OtherAssets = extract.Assets(scanner)

	print(d)
}

func print(d *declaration.Declaration) {
	fmt.Printf("Fecha: %v\nCedula: %d\nName: %s\nInstitution: %s\nJob: %s\n",
		d.Date, d.Cedula, d.Nombre+" "+d.Apellido, d.Institucion, d.Funcion)

	/*fmt.Printf("\nDepósitos:\n")
	for i, deposit := range d.Deposits {
		fmt.Println(deposit)
		if i > 1 {
			fmt.Println("...")
			break
		}
	}

	fmt.Print("\nCuentas a cobrar:\n")
	for i, debtor := range d.Debtors {
		fmt.Println(debtor)
		if i > 1 {
			fmt.Println("...")
			break
		}
	}*/

	/*fmt.Print("\nInmuebles:\n")
	for _, state := range d.RealStates {
		fmt.Println(state)
	}*/

	/*fmt.Print("\nVehículos:\n")
	for _, vehicle := range d.Vehicles {
		fmt.Println(vehicle)
	}*/

	/*fmt.Print("\nActividad Agropecuaria:\n")
	for _, activity := range d.Agricultural {
		fmt.Println(activity)
	}*/

	/*fmt.Print("\nMuebles:\n")
	for _, furnishing := range d.Furniture {
		fmt.Println(furnishing)
	}*/

	/*fmt.Print("\nOtros activos:\n")
	for _, asset := range d.OtherAssets {
		fmt.Println(asset)
	}*/
}

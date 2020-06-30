package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"code.sajari.com/docconv"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"

	"ddjj/parser/declaration"
	"ddjj/parser/extract"
)

func upload(w http.ResponseWriter, req *http.Request) {
	// 2 MB
	req.ParseMultipartForm(2 << 20)

	// in your case file would be fileupload
	file, header, err := req.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Printf("File name %s\n", header.Filename)

	err = extractPDF(file)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "No se pudo procesar el documento")

		fmt.Printf("Failed to process file %s\n", header.Filename)
		return
	}

	fmt.Fprintf(w, "Documento procesado correctamente")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/upload", upload)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func extractPDF(file io.Reader) error {
	res, err := docconv.Convert(file, "application/pdf", true)
	if err != nil {
		log.Fatal(err)
	}

	// Basic Info.
	scanner := bufio.NewScanner(strings.NewReader(res.Body))
	d := &declaration.Declaration{
		Date:        extract.Date(scanner),
		Cedula:      extract.Cedula(scanner),
		Nombre:      extract.Name(scanner),
		Apellido:    extract.Lastname(scanner),
		Institucion: extract.Institution(scanner),
		Funcion:     extract.JobTitle(scanner),
	}

	// Deposits.
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Deposits, err = extract.Deposits(scanner)
	if err != nil {
		return errors.Wrap(err, "failed when extracting deposits")
	}

	// Debtors.
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Debtors, err = extract.Debtors(scanner)
	if err != nil {
		return errors.Wrap(err, "failed when extracting debtors")
	}

	// Real state.
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.RealStates, err = extract.RealStates(scanner)
	if err != nil {
		return errors.Wrap(err, "failed when extracting debtors")
	}

	// Vehicles
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Vehicles, err = extract.Vehicles(scanner)
	if err != nil {
		return errors.Wrap(err, "failed when extracting vehicles")
	}

	// Agricultural activity
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Agricultural, err = extract.Agricultural(scanner)
	if err != nil {
		return errors.Wrap(err, "failed when extracting agricultural activities")
	}

	// Furniture
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Furniture, err = extract.Furniture(scanner)
	if err != nil {
		return errors.Wrap(err, "failed when extracting furniture")
	}

	// Other assets
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.OtherAssets, err = extract.Assets(scanner)
	if err != nil {
		return errors.Wrap(err, "failed when extracting other assets")
	}

	// Debts
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Debts, err = extract.Debts(scanner)
	if err != nil {
		return errors.Wrap(err, "failed when extracting debts")
	}

	print(d)

	scanner = bufio.NewScanner(strings.NewReader(res.Body))

	return check(scanner, d)
}

func print(d *declaration.Declaration) {
	fmt.Printf("Fecha: %v\nCedula: %d\nName: %s\nInstitution: %s\nJob: %s\n\n",
		d.Date, d.Cedula, d.Nombre+" "+d.Apellido, d.Institucion, d.Funcion)
}

func check(scanner *bufio.Scanner, d *declaration.Declaration) error {
	net := d.Net()
	scanner = extract.MoveUntil(scanner, "PATRIMONIO NETO", true)

	for i := 0; i < 6; i++ {
		scanner.Scan()
	}

	line := scanner.Text()
	if line == "" {
		return errors.New("could not get net patrimony")
	}
	expected := extract.StringToInt64(line)

	if net != expected {
		return errors.New("patrimony does not match")
	}

	return nil
}

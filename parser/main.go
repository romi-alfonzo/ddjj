package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"code.sajari.com/docconv"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gvso/ddjj/parser/database"
	"github.com/gvso/ddjj/parser/declaration"
	"github.com/gvso/ddjj/parser/extract"
)

func makeUploadHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// 2 MB
		req.ParseMultipartForm(2 << 20)

		file, header, err := req.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer file.Close()

		fmt.Printf("File name %s\n", header.Filename)

		dec, err := extractPDF(file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "No se pudo procesar el documento")

			log.Errorf("Failed to process file %s: %s\n\n", header.Filename, err)
			return
		}

		ctx := context.Background()
		d := &declaration.Declaration{}
		err = db.Collection("declarations").FindOne(ctx, map[string]interface{}{
			"cedula": dec.Cedula,
			"fecha":  dec.Fecha,
		}).Decode(d)

		if err == nil {
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(d)
			return
		}

		res, err := db.Collection("declarations").InsertOne(ctx, dec)
		if err != nil {
			fmt.Fprintf(w, "No se pudo procesar el documento")

			log.Errorf("Failed to store declaration %s: %s\n\n", header.Filename, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		dec.ID = res.InsertedID.(primitive.ObjectID)
		json.NewEncoder(w).Encode(dec)
	}
}

func handleSingleFile(filePath string) {

    //log.Println(filePath)
    dat, err := os.Open(filePath)
    if err != nil {
        log.Error("File %s not found. %s", filePath, err)
    }
    dec, err := extractPDF(dat)
    if err != nil {
            log.Errorf("Failed to process file %s: %s\n\n", filePath, err)
            return
    }

    b, err := json.Marshal(dec)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(b))
}

func main() {

        if os.Args[1] != "" {
            handleSingleFile(os.Args[1])
            return
        }

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}


	dbOpts := &database.Opts{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
		URI:      os.Getenv("DB_URI"),
	}
	db, err := database.StartConnection(dbOpts)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	http.HandleFunc("/upload", makeUploadHandler(db))

	port := os.Getenv("PORT")
	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func extractPDF(file io.Reader) (*declaration.Declaration, error) {
	res, err := docconv.Convert(file, "application/pdf", true)
	if err != nil {
		log.Fatal(err)
	}

	// Basic Info.
	scanner := bufio.NewScanner(strings.NewReader(res.Body))
	d := &declaration.Declaration{
		Fecha:       extract.Date(scanner),
		Cedula:      extract.Cedula(scanner),
		Nombre:      extract.Name(scanner),
		Apellido:    extract.Lastname(scanner),
		Institucion: extract.Institution(scanner),
		Cargo:       extract.JobTitle(scanner),
	}

	// Deposits.
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Deposits, err = extract.Deposits(scanner)
	if err != nil {
		return nil, errors.Wrap(err, "failed when extracting deposits")
	}

	// Debtors.
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Debtors, err = extract.Debtors(scanner)
	if err != nil {
		return nil, errors.Wrap(err, "failed when extracting debtors")
	}

	// Real state.
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.RealStates, err = extract.RealStates(scanner)
	if err != nil {
		return nil, errors.Wrap(err, "failed when extracting debtors")
	}

	// Vehicles
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Vehicles, err = extract.Vehicles(scanner)
	if err != nil {
		return nil, errors.Wrap(err, "failed when extracting vehicles")
	}

	// Agricultural activity
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Agricultural, err = extract.Agricultural(scanner)
	if err != nil {
		return nil, errors.Wrap(err, "failed when extracting agricultural activities")
	}

	// Furniture
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Furniture, err = extract.Furniture(scanner)
	if err != nil {
		return nil, errors.Wrap(err, "failed when extracting furniture")
	}

	// Other assets
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.OtherAssets, err = extract.Assets(scanner)
	if err != nil {
		return nil, errors.Wrap(err, "failed when extracting other assets")
	}

	// Debts
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Debts, err = extract.Debts(scanner)
	if err != nil {
		return nil, errors.Wrap(err, "failed when extracting debts")
	}

	// Income and Expenses
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.IncomeMonthly = extract.MonthlyIncome(scanner)

	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.IncomeAnnual = extract.AnnualIncome(scanner)

	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.ExpensesMonthly = extract.MonthlyExpenses(scanner)

	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.ExpensesAnnual = extract.AnnualExpenses(scanner)

	print(d)

	scanner = bufio.NewScanner(strings.NewReader(res.Body))

	err = check(scanner, d)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check correctness")
	}

	return d, nil
}

func print(d *declaration.Declaration) {
	fmt.Printf("Fecha: %v\nCedula: %d\nName: %s\nInstitution: %s\nJob: %s\n\n",
		d.Fecha, d.Cedula, d.Nombre+" "+d.Apellido, d.Institucion, d.Cargo)
}

func check(scanner *bufio.Scanner, d *declaration.Declaration) error {
	net := d.CalculatePatrimony()
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

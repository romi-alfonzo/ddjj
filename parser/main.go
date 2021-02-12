package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"code.sajari.com/docconv"

	"github.com/InstIDEA/ddjj/parser/declaration"
	"github.com/InstIDEA/ddjj/parser/extract"
)

func handleSingleFile(filePath string) extract.ParserData {
	dat, err := os.Open(filePath)
	parserData := extract.Create()
	if err == nil {
		dec, err := extractPDF(parserData, dat)
		if err != nil {
			parserData.AddMessage(fmt.Sprint("Failed to process file", filePath, ": ", err))
			parserData.Status = 1
			return parserData
		}
		parserData.SetData(dec)
	} else {
		parserData.AddMessage(fmt.Sprint("File ", filePath, " not found. ", err))
	}
	parserData.Status = 0
	return parserData
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: ./parser file.pdf")
		os.Exit(1)
		return
	}
	parsed := handleSingleFile(os.Args[1])
	parsed.ParserPrint()
}

func extractPDF(parser extract.ParserData, file io.Reader) (*declaration.Declaration, error) {
	res, err := docconv.Convert(file, "application/pdf", true)
	if err != nil {
		parser.AddMessage(err.Error())
		return nil, err
	}

	parser.RawData(res.Body)

	body := &res.Body
	d := &declaration.Declaration{}

	// Basic Info.
	scanner := bufio.NewScanner(strings.NewReader(res.Body))
	d.Fecha = parser.CheckTime(extract.Date(scanner))

	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Cedula = parser.CheckInt(extract.Cedula(scanner))

	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Nombre = parser.Check(extract.Name(scanner))

	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Apellido = parser.Check(extract.Lastname(scanner))

	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Institucion = parser.Check(extract.Institution(scanner))

	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Cargo = parser.Check(extract.JobTitle(scanner))

	// Deposits
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Deposits, err = extract.Deposits(scanner)
	if err != nil {
		parser.AddError(err)
	}

	// Debtors.
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Debtors, err = extract.Debtors(scanner)
	if err != nil {
		parser.AddError(err)
	}

	// Real state.
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.RealStates, err = extract.RealStates(scanner)
	if err != nil {
		parser.AddError(err)
	}

	// Vehicles
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Vehicles, err = extract.Vehicles(scanner)
	if err != nil {
		parser.AddError(err)
	}

	// Agricultural activity
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Agricultural, err = extract.Agricultural(scanner)
	if err != nil {
		parser.AddError(err)
	}

	// Furniture
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Furniture, err = extract.Furniture(scanner)
	if err != nil {
		parser.AddError(err)
	}

	// Other assets
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.OtherAssets, err = extract.Assets(scanner)
	if err != nil {
		parser.AddError(err)
	}

	// Debts
	scanner = bufio.NewScanner(strings.NewReader(res.Body))
	d.Debts, err = extract.Debts(scanner)
	if err != nil {
		parser.AddError(err)
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

	// Summary
	d.Resumen = extract.GetSummary(body)

	d.CalculatePatrimony()

	if d.Assets != d.Resumen.TotalActivo {
		parser.AddMessage("calculated assets and summary assets does not match")
	}

	if d.Liabilities != d.Resumen.TotalPasivo {
		parser.AddMessage("calculated liabilities and summary liabilities does not match")
	}

	if d.NetPatrimony != d.Resumen.PatrimonioNeto {
		parser.AddMessage("calculated net patrimony and summary net patrimony does not match")
	}

	return d, nil
}

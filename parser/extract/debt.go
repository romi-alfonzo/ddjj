package extract

import (
	"bufio"
	"ddjj/parser/declaration"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var totalDebt int64

var debtItemNumber int

var skipDebt = []string{
	"#",
	"TIPO DEUDA",
	"EMPRESA",
	"PLAZO",
	"CUOTA MENSUAL",
	"TOTAL DEUDA",
	"SALDO DEUDA",
}

// Debts returns money the official owes.
func Debts(scanner *bufio.Scanner) ([]*declaration.Debt, error) {

	scanner = MoveUntil(scanner, "2.1 TIPOS DE DEUDAS", true)
	var debts []*declaration.Debt

	values := [6]string{}
	index := 0
	debtItemNumber = 1

	// Also wants to skip item number
	skipDebt = append(skipDebt, strconv.Itoa(debtItemNumber))

	line, _ := getDebtLine(scanner)
	for line != "" {

		values[index] = line

		// After reading all the possible values for a single item.
		if index == 5 {
			debt := getDebt(values)

			debts = append(debts, debt)

			// Skip the next item number.
			debtItemNumber++
			skipDebt[len(skipDebt)-1] = strconv.Itoa(debtItemNumber)

			index = -1
		}

		index++

		//var nextPage bool
		line, _ = getDebtLine(scanner)
	}

	total := addDebts(debts)
	if total != totalDebt {
		for _, debt := range debts {
			fmt.Println(debt)
		}
		return nil, errors.New("The amount in debts do not match")
	}

	return debts, nil
}

func getDebt(values [6]string) *declaration.Debt {
	return &declaration.Debt{
		Tipo:    values[0],
		Empresa: values[1],
		Plazo:   stringToInt(values[2]),
		Cuota:   stringToInt64(values[3]),
		Total:   stringToInt64(values[4]),
		Saldo:   stringToInt64(values[5]),
	}
}

func getDebtLine(scanner *bufio.Scanner) (line string, nextPage bool) {
	for scanner.Scan() {
		line = scanner.Text()

		// Stop looking for debts when this is found.
		if line == "TOTALES" {
			totalDebt = getTotalInCategory(scanner)

			// Next page or end.
			scanner = MoveUntil(scanner, "TIPO DEUDA", true)
			line = scanner.Text()
			nextPage = true

			debtItemNumber = 1
			skipDebt[len(skipDebt)-1] = strconv.Itoa(debtItemNumber)
		}

		if strings.Contains(line, "OBS:") || strings.Contains(line, "RECEPCIONADO EL:") {
			continue
		}
		if isDate(line) || isBarCode(line) {
			continue
		}
		if line == "" || contains(skipDebt, line) {
			continue
		}

		return line, nextPage
	}

	return "", false
}

func addDebts(debts []*declaration.Debt) int64 {
	var total int64
	for _, d := range debts {
		total += d.Saldo
	}

	return total
}

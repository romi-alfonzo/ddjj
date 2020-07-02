package extract

import (
	"bufio"
	"fmt"
)

// MonthlyIncome extracts the monthly income.
func MonthlyIncome(scanner *bufio.Scanner) *int64 {
	scanner = MoveUntil(scanner, "TOTAL INGRESOS MENSUALES", true)
	scanner.Scan()
	scanner.Scan()
	line := scanner.Text()
	if line == "" {
		return nil
	}

	income := stringToInt64(line)
	return &income
}

// AnnualIncome extracts the annual income.
func AnnualIncome(scanner *bufio.Scanner) *int64 {
	scanner = MoveUntil(scanner, "3.2 INGRESOS ANUALES", true)
	fmt.Println("here")

	var previous string
	line := scanner.Text()
	for line != "TOTAL INGRESOS ANUALES" && scanner.Scan() {
		previous = line
		line = scanner.Text()

		for line == "" && scanner.Scan() {
			scanner.Scan()
			line = scanner.Text()
		}
	}

	fmt.Println("here2")
	income := stringToInt64(previous)
	return &income
}

// MonthlyExpenses extracts the annual income.
func MonthlyExpenses(scanner *bufio.Scanner) *int64 {
	scanner = MoveUntil(scanner, "TOTAL EGRESOS MENSUALES", true)
	scanner.Scan()
	scanner.Scan()
	line := scanner.Text()
	if line == "" {
		return nil
	}

	expense := stringToInt64(line)
	return &expense
}

// AnnualExpenses extracts the annual income.
func AnnualExpenses(scanner *bufio.Scanner) *int64 {
	scanner = MoveUntil(scanner, "3.3 EGRESOS MENSUALES", true)

	fmt.Println("hey")

	var previous string
	line := scanner.Text()
	for line != "TOTAL EGRESOS ANUALES" && scanner.Scan() {
		previous = line

		line = scanner.Text()
		for line == "" && scanner.Scan() {
			scanner.Scan()
			line = scanner.Text()
		}
	}

	fmt.Println("hey2")

	income := stringToInt64(previous)
	return &income
}

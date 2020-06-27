package extract

import (
	"bufio"
	"strconv"
	"strings"
)

// Year returns the year for the declaration.
func Year(scanner *bufio.Scanner) int {
	date := getString(scanner, "DECLARACIÓN JURADA AL :", 2)
	if date == "" {
		return 0
	}

	year := date[len(date)-4:]
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return 0
	}

	return yearInt
}

// Cedula returns the ID card number.
func Cedula(scanner *bufio.Scanner) int {
	return getInt(scanner, "CÉDULA DE IDENTIDAD:", 2)
}

// Name returns the official's name.
func Name(scanner *bufio.Scanner) string {
	return getString(scanner, "NOMBRE:", 2)
}

// Lastname returns the official's lastname.
func Lastname(scanner *bufio.Scanner) string {
	return getString(scanner, "APELLIDOS:", 2)
}

// Institution returns the official's work place.
func Institution(scanner *bufio.Scanner) string {
	return getString(scanner, "DIRECCIÓN:", 2)
}

// JobTitle returns the official's job title.
func JobTitle(scanner *bufio.Scanner) string {
	line := getString(scanner, "CARGO:", 0)

	parts := strings.Split(line, "CARGO:")

	return strings.TrimSpace(parts[1])
}

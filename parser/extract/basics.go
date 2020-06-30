package extract

import (
	"bufio"
	"strings"
	"time"
)

// Date returns the date for the declaration.
func Date(scanner *bufio.Scanner) time.Time {
	date := getString(scanner, "DECLARACIÓN JURADA AL :", 2)
	if date == "" {
		return time.Time{}
	}

	t, err := time.Parse("02/01/2006", date)
	if err != nil {
		return time.Time{}
	}

	return t
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

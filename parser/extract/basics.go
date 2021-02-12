package extract

import (
	"bufio"
	"github.com/pkg/errors"
	"time"
)

// Date returns the date for the declaration.
func Date(scanner *bufio.Scanner) (time.Time, error) {
	date := getString(scanner, "DECLARACIÓN", EVdate, nil)
	if date == "" {
		return time.Time{}, errors.New("Failed when extracting date")
	}

	t, err := time.Parse("02/01/2006", date)
	if err != nil {
		return time.Time{}, errors.New("Error parsing " + date + err.Error())
	}

	return t, nil
}

// Cedula returns the ID card number.
func Cedula(scanner *bufio.Scanner) (int, error) {
	value := getInt(scanner, "CÉDULA", EVnum, nil)
	if value == 0 {
		return 0, errors.New("failed when extracting cedula")
	}
	return value, nil
}

// Name returns the official's name.
func Name(scanner *bufio.Scanner) (string, error) {
	value := getString(scanner, "NOMBRE", EValphaNum, nil)
	if value == "" {
		return "", errors.New("failed when extracting name")
	}
	return value, nil
}

// Lastname returns the official's lastname.
func Lastname(scanner *bufio.Scanner) (string, error) {
	value := getString(scanner, "APELLIDOS", EValphaNum, nil)
	if value == "" {
		return "", errors.New("failed when extracting lastname")
	}
	return value, nil
}

// Institution returns the official's work place.
func Institution(scanner *bufio.Scanner) (string, error) {
	value := getString(scanner, "DIRECCIÓN", EValphaNum, nil)
	if value == "" {
		return "", errors.New("failed when extracting institucion")
	}
	return value, nil
}

// JobTitle returns the official's job title.
func JobTitle(scanner *bufio.Scanner) (string, error) {
	value := getString(scanner, "CARGO", EValphaNum, nil)
	if value == "" {
		return "", errors.New("failed when extracting cargo")
	}
	return value, nil
}

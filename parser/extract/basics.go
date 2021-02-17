package extract

import (
	"github.com/pkg/errors"
	"time"
)

// Date returns the date for the declaration.
func Date(e *Extractor) (time.Time, error) {
	var date string

	e.BindFlag(EXTRACTOR_FLAG_1)
	if e.MoveUntilContains(CurrToken, "DECLARACIÓN") {
		for e.Scan() {
			if isDate(e.CurrToken) {
				date = e.CurrToken
				break
			}
		}
	}

	if date == "" {
		return time.Time{}, errors.New("failed when extracting date")
	}

	t, err := time.Parse("02/01/2006", date)
	if err != nil {
		return time.Time{}, errors.New("Error parsing " + date + err.Error())
	}
	return t, nil
}

// Cedula returns the ID card number.
func Cedula(e *Extractor) (int, error) {
	var value int

	e.BindFlag(EXTRACTOR_FLAG_1)
	if e.MoveUntilStartWith(CurrToken, "CÉDULA") {
		if isNumber(e.NextToken) {
			value = stringToInt(e.NextToken)
		}
	}

	if value == 0 {
		return 0, errors.New("failed when extracting cedula")
	}
	return value, nil
}

// Name returns the official's name.
func Name(e *Extractor) (string, error) {
	var value string

	e.BindFlag(EXTRACTOR_FLAG_1)
	if e.MoveUntilStartWith(CurrToken, "NOMBRE") {
		if isAlpha(e.NextToken) {
			value = e.NextToken
		}
	}

	if value == "" {
		return "", errors.New("failed when extracting name")
	}
	return value, nil
}

// Lastname returns the official's lastname.
func Lastname(e *Extractor) (string, error) {
	var value string

	e.BindFlag(EXTRACTOR_FLAG_1)
	if e.MoveUntilStartWith(CurrToken, "APELLIDOS") {
		if isAlpha(e.NextToken) {
			value = e.NextToken
		}
	}

	if value == "" {
		return "", errors.New("failed when extracting lastname")
	}
	return value, nil
}

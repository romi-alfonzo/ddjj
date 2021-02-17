package extract

import (
	"fmt"
	"strings"
	"github.com/InstIDEA/ddjj/parser/declaration"	
)

// TODO
// validate the institution with known institutions
// to reduce data type checking and content controls

func Jobs(e *Extractor, parser *ParserData) ([]*declaration.Job) {

	var instituciones []*declaration.Job
	var resultsPositions []int
	var counter = countJobs(e)
	var successful int

	// the institution/job name can be ignored
	// if contains any element from the list
	exclude := []string {
		"TIPO",
		"DEPENDENCIA",	
		"CONTRATADO",
		"ACTO ADMINIST.",
		"COMISIONADO",
		"CATEGORÍA",
		"TELÉFONO",
		"ACTO ADM. COM.",
		"FECHA INGRESO",
		"FECHA EGRESO",
		"PRINCIPAL",
		"DECRETO",
		"NOMBRADO",
		"NOMBRAMIENTO",
		"FECHA ASUNC",
	}

	e.Rewind()
	e.BindFlag(EXTRACTOR_FLAG_1)

	if counter > 0 &&
	e.MoveUntilStartWith(CurrToken, "DATOS LABORALES") {
		e.SaveLine()
		job := &declaration.Job{ }

		for e.Scan() {
			if counter == successful {
				break
			}

			if job.Cargo == "" {
				job.Cargo = getJobTitle(e, exclude, &resultsPositions)
			}

			if job.Cargo != "" &&
			job.Institucion == "" {
				job.Institucion = getJobInstitution(e, exclude, &resultsPositions)
			}

			if job.Cargo != "" && job.Institucion != "" {
				successful++
				instituciones = append(instituciones, job)
				job = &declaration.Job{ }
				e.MoveUntilSavedLine()
			}
		}
	}

	if successful != counter {
		parser.addMessage(fmt.Sprintf("ignored jobs: %d/%d", counter - successful, counter))
	}

	if instituciones == nil {
		parser.addError(fmt.Errorf("failed when extracting jobs"))
		return nil
	}

	return instituciones
}

func getJobTitle(e *Extractor, exclude []string, pos *[]int) string {
	// control positions are based on first matching tokens

	var value string

	if strings.Contains(e.CurrToken, "CARGO") &&
	!ContainsIntItem(*pos, e.CurrLineNum()) {
		val, check := isKeyValuePair(e.CurrToken, "CARGO")
		if check {
			*pos = append(*pos, e.CurrLineNum())
			e.MoveUntilSavedLine()
			return val
		}
	}

	if isCurrLine(e.CurrToken, "INSTITUCIÓN") {
		if !ContainsItem(exclude, e.PrevToken) &&
		!ContainsIntItem(*pos, e.PrevLineNum()) &&
		!isCurrLine(e.PrevToken, "SI") &&
		!isNumber(e.PrevToken) {
			value = e.PrevToken
			*pos = append(*pos, e.PrevLineNum())
			e.MoveUntilSavedLine()
			return value
		}

		if !ContainsItem(exclude, e.NextToken) &&
		!ContainsIntItem(*pos, e.NextLineNum()) {
			value = e.NextToken
			*pos = append(*pos, e.NextLineNum())
			e.MoveUntilSavedLine()
			return value
		}
	}

	if isCurrLine(e.PrevToken, "TELÉFONO") &&
	!isCurrLine(e.NextToken, "CARGO") &&
	!ContainsItem(exclude, e.CurrToken) &&
	!ContainsIntItem(*pos, e.CurrLineNum()) &&
	len(e.CurrToken) > 5 && // minimum length for job title
	!isNumber(e.CurrToken) &&
	!isDate(e.CurrToken) {
		value = e.CurrToken
		*pos = append(*pos, e.CurrLineNum())
		e.MoveUntilSavedLine()
		return value
	}

	return ""
}

func getJobInstitution(e *Extractor, exclude []string, pos *[]int) string {
	// control positions are based on first matching tokens

	var value string

	if isCurrLine(e.PrevToken, "INSTITUCIÓN") &&
	isCurrLine(e.CurrToken, "TIPO") &&
	!ContainsItem(exclude, e.NextToken) &&
	!ContainsIntItem(*pos, e.NextLineNum()) &&
	!strings.Contains(e.NextToken, "DIRECC") &&
	len(e.NextToken) > 5 { // minimum length for institute
		value = e.NextToken
		*pos = append(*pos, e.NextLineNum())
		e.MoveUntilSavedLine()
		return value
	}

	if isCurrLine(e.PrevToken, "DIRECCIÓN") {
		_, check := isKeyValuePair(e.PrevToken, "DIRECCIÓN")

		if check &&
		!ContainsItem(exclude, e.NextToken) &&
		!ContainsIntItem(*pos, e.NextLineNum()) &&
		!isDate(e.NextToken) {
			value = e.NextToken
			*pos = append(*pos, e.NextLineNum())
			e.MoveUntilSavedLine()
			return value
		}

		if !ContainsItem(exclude, e.CurrToken) &&
		!ContainsIntItem(*pos, e.CurrLineNum()) {
			value = e.CurrToken
			*pos = append(*pos, e.CurrLineNum())
			e.MoveUntilSavedLine()
			return value
		}
	}

	if isCurrLine(e.CurrToken, "FECHA ASUNC") &&
	!ContainsItem(exclude, e.NextToken) &&
	!ContainsIntItem(*pos, e.NextLineNum()) &&
	len(e.NextToken) > 5 && // minimum length for institute
	!isDate(e.NextToken) {
		value = e.NextToken
		*pos = append(*pos, e.NextLineNum())
		e.MoveUntilSavedLine()
		return value
	}

	return ""
}

func countJobs(e *Extractor) int {
	var counter int
	for e.Scan() {
		if isCurrLine(e.NextToken, "página") {
			break
		}

		if strings.Contains(e.CurrToken, "CARGO:") {
			counter++
		}
	}
	return counter
}

package extract

import (
	"fmt"
	"strings"

	"github.com/InstIDEA/ddjj/parser/declaration"
)

// Debtors returns the debts people have with the official.
func Debtors(e *Extractor, parser *ParserData) ([]*declaration.Debtor, error) {
	var debtors []*declaration.Debtor //lsit of extracted debtors
	debt := &declaration.Debtor{}     //aux for the actual extraction
	e.BindFlag(EXTRACTOR_FLAG_1)      //remueve las lineas en blanco
	e.BindFlag(EXTRACTOR_FLAG_2)      //remueve los espacios en los extremos
	//EXTRACTOR_FLAG_3 crea nuevos tokens siempre que dentro de la linea haya mas o igual a 3 espacios
	var bandera bool
	bandera = false
	counter := 0
	successful := 0
	if e.MoveUntilStartWith(CurrToken, "1.3 CUENTAS A COBRAR") {
		for e.Scan() {
			// other assets table header and OBS are omitted
			if isAssetFormField(e.CurrToken) {
				bandera = true //we are in the table records because we have the header
				continue
			}
			if strings.Contains(e.CurrToken, "OBS:") && bandera {
				counter++
				continue
			}
			// final of others assets of current page
			if strings.Contains(e.CurrToken, "TOTAL CUENTAS POR COBRAR:") {
				bandera = false
			}
			//if the ban it's true, we can proceed with the extraction
			if bandera {
				values := tokenize(e.CurrToken, 3)
				if len(values) == 5 {
					debt = detDebtor(values)
					debtors = append(debtors, debt)
				}
			}
		}
		successful = len(debtors)
	}
	if successful != counter {
		parser.addMessage(fmt.Sprintf("ignored debtors: %d/%d", counter-successful, counter))
	}

	if debtors == nil {
		parser.addError(fmt.Errorf("failed when extracting debtors"))
		return nil, nil
	}

	return debtors, nil
}

/*
Function to check if a given string is or not the header of the section.
Parameter: string s
Return: True or false
*/

func isAssetFormField(s string) bool {
	formField := []string{
		"#",
		"NOMBRE DEL DEUDOR",
		"CLASE (A LA VISTA O PLAZOS)",
		"PLAZO EN",
		"IMPORTE",
	}

	s = removeAccents(s)
	for _, value := range formField {
		if !strings.Contains(s, value) {
			return false
		}
	}

	return true
}

/*
Function to load the extracted values into the OtherAsset structure.
Parameters: values in an array of strings. The first element is not inserted because it is the index and not relevant.
Return: an instance of OtherAsset with the values from the array
*/

func detDebtor(values []string) *declaration.Debtor {
	return &declaration.Debtor{
		Nombre:  values[1],
		Clase:   values[2],
		Plazo:   stringToInt(values[3]),
		Importe: stringToInt64(values[4]),
	}
}

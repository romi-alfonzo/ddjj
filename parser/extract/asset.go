package extract

import (
	"fmt"
	"strings"

	"github.com/InstIDEA/ddjj/parser/declaration"
)

func Assets(e *Extractor, parser *ParserData) ([]*declaration.OtherAsset, error) {
	var assets []*declaration.OtherAsset //lsit of extracted assets
	asset := &declaration.OtherAsset{}   //aux for the actual extraction
	e.BindFlag(EXTRACTOR_FLAG_1)         //remueve las lineas en blanco
	e.BindFlag(EXTRACTOR_FLAG_2)         //remueve los espacios en los extremos
	//EXTRACTOR_FLAG_3 crea nuevos tokens siempre que dentro de la linea haya mas o igual a 3 espacios
	var bandera bool
	bandera = false
	counter := 0
	successful := 0
	if e.MoveUntilStartWith(CurrToken, "1.9 OTROS ACTIVOS") {
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
			if strings.Contains(e.CurrToken, "TOTAL OTROS ACTIVOS") {
				bandera = false
			}
			//if the ban it's true, we can proceed with the extraction
			if bandera {
				values := tokenize(e.CurrToken, 3)
				//case 1: Description is in two lines
				//in this case the lines are
				//descPart1
				//number of the register
				//descPart2
				//rest of row
				if len(values) == 1 && isNumber(e.CurrToken) {
					description := e.PrevToken + " " + e.NextToken
					// moving the current token to the next part
					e.Scan()
					e.Scan()

					//building the struct of other assets
					fixed := []string{"#", description}
					values = append(fixed, tokenize(e.CurrToken, 3)...)
				} else
				//case 2: Enterprise name is in two lines
				//in this case the lines are
				//enterprisePart1
				//number of the register + description
				//enterprisePart2
				//rest of row
				if len(values) == 2 {
					enterpriseNamePart1 := e.PrevToken
					//extracting the description of the currentToken thats saved on values array
					description := values[1]
					e.Scan() // we need to save the description in this part
					allName := enterpriseNamePart1 + " " + e.CurrToken
					//moving to the rest of the row
					e.Scan()

					//building the struct of other assets
					fixed := []string{"#", description, allName}
					values = append(fixed, tokenize(e.CurrToken, 3)...)

				} else
				//case 3: country in two lines
				//namePart1
				//num + description + enterprise + ruc
				//namePart2
				//cant + price + total
				if len(values) == 4 {
					country := e.PrevToken + " " + e.NextToken
					description := values[1]
					enterprise := values[2]
					ruc := values[3]
					// moving the current token to the next part
					e.Scan()
					e.Scan()

					//building the struct of other assets
					fixed := []string{"#", description, enterprise, ruc, country}
					values = append(fixed, tokenize(e.CurrToken, 4)...)
				}

				if len(values) == 8 {
					asset = getAsset(values)
					assets = append(assets, asset)
				}
			}
		}
		successful = len(assets)
	}
	if successful != counter {
		fmt.Println(successful)
		parser.addMessage(fmt.Sprintf("ignored assets: %d/%d", counter-successful, counter))
	}

	if assets == nil {
		parser.addError(fmt.Errorf("failed when extracting assets"))
		return nil, nil
	}

	return assets, nil
}

/*
Function to check if a given string is or not the header of the section.
Parameter: string s
Return: True or false
*/

func isAssetFormField(s string) bool {
	formField := []string{
		"DESCRIPCION",
		"EMPRESA",
		"RUC",
		"PAIS",
		"CANT.",
		"PRECIO UNI.",
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

func getAsset(values []string) *declaration.OtherAsset {
	return &declaration.OtherAsset{
		Descripcion: values[1],
		Empresa:     values[2],
		RUC:         values[3],
		Pais:        values[4],
		Cantidad:    stringToInt64(values[5]),
		Precio:      stringToInt64(values[6]),
		Importe:     stringToInt64(values[7]),
	}
}

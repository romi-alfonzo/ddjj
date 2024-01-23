package extract

import (
	"fmt"
	"strings"

	"github.com/InstIDEA/ddjj/parser/declaration"
)

func Assets(e *Extractor, parser *ParserData) ([]*declaration.OtherAsset, error) {

	e.BindFlag(EXTRACTOR_FLAG_1) //remueve las lineas en blanco
	e.BindFlag(EXTRACTOR_FLAG_2) //remueve los espacios en los extremos
	//EXTRACTOR_FLAG_3 crea nuevos tokens siempre que dentro de la linea haya mas o igual a 3 espacios

	if e.MoveUntilStartWith(CurrToken, "#              DESCRIPCIÓN") {
		for e.Scan() {
			fmt.Println(e.CurrToken)
			if strings.Contains(e.CurrToken, "TOTAL OTROS ACTIVOS") {
				break
			}
		}
	}
	return nil, nil
}

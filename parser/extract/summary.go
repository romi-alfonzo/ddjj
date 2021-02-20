package extract

import (
	"github.com/InstIDEA/ddjj/parser/declaration"
	"fmt"
)

func Summary(e *Extractor, parser *ParserData) *declaration.Summary {
	var index int

	results := [3]int64{ -1, -1, -1 }
	e.BindFlag(EXTRACTOR_FLAG_1)

	if e.MoveUntilContains(CurrToken, "RESUMEN") {
		for e.Scan() {
			if index > 2 {
				break
			}

			if isNumber(e.CurrToken) {
				results[index] = StringToInt64(e.CurrToken)
				index++
			}
		}
	}

	if results[0] == -1 &&
	results[1] == -1 &&
	results[2] == -1 {
		parser.addError(fmt.Errorf("failed when extracting summary"))
		return nil
	}

	return &declaration.Summary{
		TotalActivo: results[0],
		TotalPasivo: results[1],
		PatrimonioNeto: results[2],
	}
}

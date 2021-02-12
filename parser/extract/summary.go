package extract

import (
	"bufio"
	"github.com/InstIDEA/ddjj/parser/declaration"
	"strings"
)

func GetSummary(body *string) *declaration.Summary {
	r := &declaration.Summary{}
	exclude := &[]int{}

	scanner := bufio.NewScanner(strings.NewReader(*body))
	r.TotalActivo = StringToInt64(getString(scanner, "TOTAL ACTIVO", EVnum, exclude))

	scanner = bufio.NewScanner(strings.NewReader(*body))
	r.TotalPasivo = StringToInt64(getString(scanner, "TOTAL PASIVO", EVnum, exclude))

	scanner = bufio.NewScanner(strings.NewReader(*body))
	r.PatrimonioNeto = StringToInt64(getString(scanner, "PATRIMONIO NETO", EVnum, exclude))

	return r
}

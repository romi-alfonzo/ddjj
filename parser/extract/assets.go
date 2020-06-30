package extract

import (
	"bufio"
	"ddjj/parser/declaration"
	"strconv"
	"strings"
)

var totalAssets int64

var assetsItemNumber int

var skipAssets = []string{
	"#",
	"DESCRIPCIÓN",
	"EMPRESA",
	"RUC",
	"PAÍS",
	"CANT.",
	"PRECIO UNI.",
	"IMPORTE",
}

// Assets returns other assets owned by the official.
func Assets(scanner *bufio.Scanner) []*declaration.OtherAsset {

	scanner = moveUntil(scanner, "1.9 OTROS ACTIVOS", true)
	var assets []*declaration.OtherAsset

	values := [7]string{}
	index := 0
	assetsItemNumber = 1

	// Also wants to skip item number
	skipAssets = append(skipAssets, strconv.Itoa(assetsItemNumber))

	line, _ := getAssetLine(scanner)
	for line != "" {

		values[index] = line

		// After reading all the possible values for a single item.
		if index == 6 {
			asset := getAsset(values)

			assets = append(assets, asset)

			// Skip the next item number.
			assetsItemNumber++
			skipAssets[len(skipAssets)-1] = strconv.Itoa(assetsItemNumber)

			index = -1
		}

		index++

		//var nextPage bool
		line, _ = getAssetLine(scanner)
	}

	/*total := addAssets(assets)
	if total != totalAssets {
		log.Fatal("The amounts in other assets do not match")
	}*/

	return assets
}

func getAsset(values [7]string) *declaration.OtherAsset {
	return &declaration.OtherAsset{
		Descripcion: values[0],
		Empresa:     values[1],
		RUC:         values[2],
		Pais:        values[3],
		Cantidad:    stringToInt64(values[4]),
		Precio:      stringToInt64(values[5]),
		Importe:     stringToInt64(values[6]),
	}
}

func getAssetLine(scanner *bufio.Scanner) (line string, nextPage bool) {
	for scanner.Scan() {
		line = scanner.Text()

		// Stop looking for assets when this is found.
		if line == "TOTAL OTROS ACTIVOS" {
			totalAssets = getTotalInCategory(scanner)

			// Next page or end.
			scanner = moveUntil(scanner, "TIPO MUEBLES", true)
			line = scanner.Text()
			nextPage = true

			assetsItemNumber = 1
			skipAssets[len(skipAssets)-1] = strconv.Itoa(assetsItemNumber)
		}

		if strings.Contains(line, "OBS:") || strings.Contains(line, "RECEPCIONADO EL:") {
			continue
		}
		if isDate(line) || isBarCode(line) {
			continue
		}
		if line == "" || contains(skipAssets, line) {
			continue
		}

		return line, nextPage
	}

	return "", false
}

func addAssets(assets []*declaration.OtherAsset) int64 {
	var total int64
	for _, a := range assets {
		total += a.Importe
	}

	return total
}

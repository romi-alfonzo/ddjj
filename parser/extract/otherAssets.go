package extract

import (
	"strings"

	"github.com/InstIDEA/ddjj/parser/declaration"
	"github.com/pkg/errors"
)

func otherAssets(e *Extractor, parser *ParserData) ([]*declaration.OtherAsset, error) {

	e.BindFlag(EXTRACTOR_FLAG_1)
	e.BindFlag(EXTRACTOR_FLAG_2)

	var assets []*declaration.OtherAsset
	assets = countAssets(e, assets)

	total := addAssets(assets)
	if total == 0 {
		return nil, errors.New("failed when extracting other assets")
	}

	return assets, nil
}

func countAssets(e *Extractor, assets []*declaration.OtherAsset) []*declaration.OtherAsset {
	asset := &declaration.OtherAsset{}
	for e.Scan() {
		if strings.Contains(e.CurrToken, "ACCIONES") {
			//Cuando el nombre de la empresa tiene dos lineas, queda en PrevToken, sino queda OBS N/A o el form field y es el caso "normal"
			if strings.Contains(e.PrevToken, "OBS: N/A") || isAssetFormField(e.PrevToken) {
				values := tokenize(e.CurrToken, 5)
				//asset is added only if it has all of the needed values
				if len(values) == 8 {
					asset = getAsset3(values)
					assets = append(assets, asset)
				} else {
					continue
				}
			} else {
				//Cuando hay dos lineas, la primera linea queda en PrevToken, en CurrToken queda el indice y ACCIONES y en NextToken la segunda linea del nombre
				//entonces se concatenan PrevToken y NextToken, y luego se vuelve a scanear para tener el resto de los datos
				//Tambien se tiene el caso en el que el nombre se encuentra una linea mas arriba a pesar de no tener dos lineas (Ejemplo: Cartes 2018, Consignataria de Ganado S.A>)
				var name string
				if !strings.Contains(e.NextToken, "OBS: N/A") {
					name = e.PrevToken + " " + e.NextToken
				}
				for i := 1; i < 3; i++ {
					e.Scan()
				}
				//aditional values that are not in the line but are needed to have the full asset, included the name
				additional := []string{"#", "ACCIONES", name}
				values := append(additional, tokenize(e.CurrToken, 4)...)
				if len(values) == 8 {
					asset = getAsset3(values)
					assets = append(assets, asset)
				} else {
					continue
				}

			}
		} else if strings.Contains(e.CurrToken, "CERTIFICADO DE DEPOSITOS DE") {
			//subsequent scans are needed due to the document format
			for i := 1; i < 4; i++ {
				e.Scan()
			}
			//fixed values that are not in the line but are needed to have the full asset
			fixed := []string{"#", "CERTIFICADO DE DEPOSITOS DE AHORROS"}
			values := append(fixed, tokenize(e.CurrToken, 4)...)
			asset = getAsset3(values)
			assets = append(assets, asset)
		} else if strings.Contains(e.CurrToken, "INVERSIONES") || strings.Contains(e.CurrToken, "BONOS") || strings.Contains(e.CurrToken, "PATENTES") || (strings.Contains(e.CurrToken, "OTROS") && strings.Contains(e.NextToken, "OBS: N/A")) {
			values := tokenize(e.CurrToken, 5)
			if len(values) == 8 {
				asset = getAsset3(values)
				assets = append(assets, asset)
			} else {
				continue
			}
		} else {
			continue
		}
	}

	return assets
}

/*
Function to calculate the total of the extracted assets.
*/

func addAssets(assets []*declaration.OtherAsset) int64 {
	var total int64
	for _, a := range assets {
		total += a.Importe
	}
	return total
}

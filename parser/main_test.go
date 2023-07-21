package main

import (
	"fmt"
	"github.com/InstIDEA/ddjj/parser/extract"
	"reflect"
	"testing"
)

func TestDarioRamon(t *testing.T) {

	data := handleSingleFile("./test_declarations/4736335_DARIO_RAMON_VERA.pdf")

	if data.Data == nil {
		t.Errorf("Error parsing the document")
	}

	for _, item := range data.Message {
		fmt.Println(item)
	}

	fmt.Printf("\n\n")
	fmt.Println("Message: ", data.Message)
	fmt.Println("Structured: ", data.Structured)

	AssertEqual(t, "DARIO RAMON", data.Data.Nombre)
	AssertEqual(t, "2016-10-19", data.Data.Fecha.Format("2006-01-02"))
	AssertEqual(t, "MILITAR", data.Data.Instituciones[0].Cargo)
	AssertEqual(t, "COMANDO DE LAS FUERZAS MILITARES", data.Data.Instituciones[0].Institucion)
	AssertEqual(t, "", data.Data.Conyuge)
	AssertEqual(t, int64(0), data.Data.Resumen.TotalActivo)
	AssertEqual(t, int64(6560000), data.Data.Resumen.TotalPasivo)
	AssertEqual(t, int64(-6560000), data.Data.Resumen.PatrimonioNeto)
}

func TestMarioAbdo2016(t *testing.T) {

	data := handleSingleFile("./test_declarations/267948_MARIO_ABDO_BENITEZ.pdf")

	if data.Data == nil {
		t.Errorf("Error parsing the document")
	}

	for _, item := range data.Message {
		fmt.Println(item)
	}

	fmt.Printf("\n\n")
	fmt.Println("Nombre: ", data.Data.Nombre)
	fmt.Println("Fecha: ", data.Data.Fecha)
	fmt.Println("Conyuge: ", data.Data.Conyuge)
	fmt.Println("Cargo: ", data.Data.Instituciones[0].Cargo)
	fmt.Println("Institucion: ", data.Data.Instituciones[0].Institucion)
	fmt.Println("Resumen Activos: ", data.Data.Resumen.TotalActivo)
	fmt.Println("Resumen Pasivos: ", data.Data.Resumen.TotalPasivo)
	fmt.Println("Resumen Patrimonio Neto: ", data.Data.Resumen.PatrimonioNeto)

	AssertEqual(t, "MARIO", data.Data.Nombre)
	AssertEqual(t, "2016-07-13", data.Data.Fecha.Format("2006-01-02"))
	AssertEqual(t, "SENADOR NACIONAL", data.Data.Instituciones[0].Cargo)
	AssertEqual(t, "CONGRESO NACIONAL", data.Data.Instituciones[0].Institucion)
	AssertEqual(t, "SILVANA MARIA AUXILIADORA LOPEZ MOREIRA BO", data.Data.Conyuge)
	AssertEqual(t, int64(3263852172), data.Data.Resumen.TotalActivo)
	AssertEqual(t, int64(241094919), data.Data.Resumen.TotalPasivo)
	AssertEqual(t, int64(3022757253), data.Data.Resumen.PatrimonioNeto)
}

func TestFatimaMagdalenaBaez2015(t *testing.T) {

	// the go parser crashing with this file
	// last debug line printed before crashing
	// [CREDITOS COOP NAZARET 30 1,200,000 30,000,000 30,000,000]

	data := handleSingleFile("./test_declarations/772198_FATIMA_MAGDALENA_BAEZ.pdf")

	if data.Data == nil {
		t.Errorf("Error parsing the document")
	}

	for _, item := range data.Message {
		fmt.Println(item)
	}

	data.Print()

	AssertEqual(t, "FATIMA MAGDALENA", data.Data.Nombre)
	AssertEqual(t, "2015-05-07", data.Data.Fecha.Format("2006-01-02"))
	//AssertEqualWM(t, "The length of the debts is incorrect", 6, len(data.Data.Debts))
	AssertEqual(t, int64(29500000), data.Data.Resumen.TotalActivo)
	AssertEqual(t, int64(278178670), data.Data.Resumen.TotalPasivo)
	AssertEqual(t, int64(-248678670), data.Data.Resumen.PatrimonioNeto)
}

func TestVictorBlancoSilva2015(t *testing.T) {

	// the go parser crashing with this file
	// doesn't get the correctly data from vehicle model "VOLKSWAGEN"
	// last debug line printed before crashing
	// Modelo: AÑO ADQUIS.: 2014
	// Fabricacion: 0

	data := handleSingleFile("./test_declarations/775679_VICTOR_BLANCO_SILVA.pdf")

	if data.Data == nil {
		t.Errorf("Error parsing the document")
	}

	for _, item := range data.Message {
		fmt.Println(item)
	}

	data.Print()

	AssertEqual(t, "VICTOR", data.Data.Nombre)
	AssertEqual(t, "2015-11-27", data.Data.Fecha.Format("2006-01-02"))
	//AssertEqual(t, int64(63000000), data.Data.Vehicles[1].Importe)
	AssertEqual(t, int64(383500000), data.Data.Resumen.TotalActivo)
	AssertEqual(t, int64(0), data.Data.Resumen.TotalPasivo)
	AssertEqual(t, int64(383500000), data.Data.Resumen.PatrimonioNeto)
}

func TestMariaLorenaRiverosMiranda2015(t *testing.T) {

	// the go parser crashing with this file
	// doesn't get the correctly data
	// last debug line printed before crashing
	// [BRISTOL S.A 12 189,000 2,268,000 849,000 11,421,000]
	data := handleSingleFile("./test_declarations/776388_MARIA_LORENA_RIVEROS_MIRANDA.pdf")

	if data.Data == nil {
		t.Errorf("Error parsing the document")
	}

	for _, item := range data.Message {
		fmt.Println(item)
	}

	// TODO fix parsing of debts
	AssertHasError(t, &data, "the amount in debts do not match (calculated=17100000 in pdf: 849000)")

	data.Print()

	AssertEqual(t, "MARIA LORENA", data.Data.Nombre)
	AssertEqual(t, "2015-03-10", data.Data.Fecha.Format("2006-01-02"))

	// this is a case of wrong active/passive/nw
	AssertEqual(t, int64(0), data.Data.Resumen.TotalActivo)
	AssertEqual(t, int64(0), data.Data.Resumen.TotalPasivo)
	AssertEqual(t, int64(0), data.Data.Resumen.PatrimonioNeto)

	//AssertEqual(t, int64(189000), data.Data.Debts[2].Cuota)
	//AssertEqual(t, int64(30000000), data.Data.Debts[2].Total)
	//AssertEqual(t, int64(30000000), data.Data.Debts[2].Saldo)
}

func TestLilianSamaniego2016(t *testing.T) {

	// program freeze after input
	// with this version and 1.0.0
	// https://github.com/InstIDEA/ddjj/releases/tag/1.0.0

	// with previus version return zero values
	// https://github.com/Ravf95/ddjj/tree/feature/local_mode/parser

	data := handleSingleFile("./test_declarations/78832_LILIAN_MARLENE_SAMANIEGO_BENEGA.pdf")

	if data.Data == nil {
		t.Errorf("Error parsing the document")
	}

	data.Print()

	AssertEqual(t, "LILIAN MARLENE", data.Data.Nombre)
	AssertEqual(t, "2016-10-07", data.Data.Fecha.Format("2006-01-02"))
}

func TestNataliaDure2019(t *testing.T) {
	// program freeze after input
	// with this version and 1.0.0
	// https://github.com/InstIDEA/ddjj/releases/tag/1.0.0

	// with previus version return zero values
	// https://github.com/Ravf95/ddjj/tree/feature/local_mode/parser

	data := handleSingleFile("./test_declarations/592859_NATALIA_ELIZABETH_DURE_CARDOZO.pdf")

	if data.Data == nil {
		t.Errorf("Error parsing the document")
	}

	data.Print()

	AssertEqual(t, "NATALIA ELIZABETH", data.Data.Nombre)
	AssertEqual(t, "2019-03-07", data.Data.Fecha.Format("2006-01-02"))
}

// AssertEqual checks if values are equal
func AssertEqual(t *testing.T, want interface{}, got interface{}) {
	if want == got {
		return
	}
	// debug.PrintStack()
	t.Errorf("Received %v (type %v), expected %v (type %v)", got, reflect.TypeOf(got), want, reflect.TypeOf(want))
}

func AssertEqualWM(t *testing.T, baseMsg string, want interface{}, got interface{}) {
	if want == got {
		return
	}
	// debug.PrintStack()
	t.Errorf(baseMsg+". Received %v (type %v), expected %v (type %v)", got, reflect.TypeOf(got), want, reflect.TypeOf(want))
}

func AssertTrue(t *testing.T, message string, toCheck bool) {
	if toCheck {
		return
	}
	// debug.PrintStack()
	t.Errorf(message)
}

func AssertHasError(t *testing.T, dat *extract.ParserData, desiredError string) {
	var found = false
	for _, row := range dat.Message {
		if row == desiredError {
			found = true
		}
	}
	AssertTrue(t, "The message: {"+desiredError+"} should be present in the errors", found)
}

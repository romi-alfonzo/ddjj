package main

import (
	"fmt"
	"reflect"
	"testing"
)

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
	fmt.Println("Resumen Activos: ", data.Data.Resumen.TotalActivo)
	fmt.Println("Resumen Pasivos: ", data.Data.Resumen.TotalPasivo)
	fmt.Println("Resumen Patrimonio Neto: ", data.Data.Resumen.PatrimonioNeto)

	AssertEqual(t, "MARIO", data.Data.Nombre)
	AssertEqual(t, "2016-07-13", data.Data.Fecha.Format("2006-01-02"))
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

	AssertEqual(t, "FATIMA", data.Data.Nombre)
	AssertEqual(t, "2015-04-14", data.Data.Fecha.Format("2006-01-02"))
	AssertEqual(t, int64(1200000), data.Data.Debts[2].Cuota)
	AssertEqual(t, int64(30000000), data.Data.Debts[2].Total)
	AssertEqual(t, int64(30000000), data.Data.Debts[2].Saldo)
}

func TestVictorBlancoSilva2015(t* testing.T) {

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
	
	AssertEqual(t, "VICTOR", data.Data.Nombre)
	AssertEqual(t, "2015-11-27", data.Data.Fecha.Format("2006-01-02"))
	AssertEqual(t, int64(63000000), data.Data.Vehicles[1].Importe)
}

func TestMariaLorenaRiverosMiranda2015(t* testing.T) {

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
	
	AssertEqual(t, "MARIA LORENA", data.Data.Nombre)
	AssertEqual(t, "2015-03-10", data.Data.Fecha.Format("2006-01-02"))
	
	AssertEqual(t, int64(189000), data.Data.Debts[2].Cuota)
	AssertEqual(t, int64(30000000), data.Data.Debts[2].Total)
	AssertEqual(t, int64(30000000), data.Data.Debts[2].Saldo)
}

// AssertEqual checks if values are equal
func AssertEqual(t *testing.T, want interface{}, got interface{}) {
	if want == got {
		return
	}
	// debug.PrintStack()
	t.Errorf("Received %v (type %v), expected %v (type %v)", got, reflect.TypeOf(got), want, reflect.TypeOf(want))
}
